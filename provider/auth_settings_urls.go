// Copyright (c) 2023 Zscaler Technology Alliances, <zscaler-partner-labs@z-bd.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package provider implements the Auth Settings URLs resource.
// Adopted from terraform-provider-zia resource_zia_auth_settings_urls.go.

package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/user_authentication_settings"
)

const authSettingsUrlsID = "all_urls"

// AuthSettingsUrls implements the zia:index:AuthSettingsUrls resource.
// Manages URLs exempted from user authentication. Singleton. Delete is a no-op.
type AuthSettingsUrls struct{}

// AuthSettingsUrlsArgs are the inputs.
type AuthSettingsUrlsArgs struct {
	Urls []string `pulumi:"urls,optional"`
}

// AuthSettingsUrlsState is the persisted state.
type AuthSettingsUrlsState struct {
	AuthSettingsUrlsArgs
	ResourceId string `pulumi:"resourceId"`
}

func (AuthSettingsUrls) Create(ctx context.Context, req infer.CreateRequest[AuthSettingsUrlsArgs]) (infer.CreateResponse[AuthSettingsUrlsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[AuthSettingsUrlsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	urls := user_authentication_settings.ExemptedUrls{URLs: req.Inputs.Urls}
	if _, err := user_authentication_settings.Update(ctx, service, urls); err != nil {
		return infer.CreateResponse[AuthSettingsUrlsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[AuthSettingsUrlsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	state := AuthSettingsUrlsState{
		AuthSettingsUrlsArgs: req.Inputs,
		ResourceId:           authSettingsUrlsID,
	}
	return infer.CreateResponse[AuthSettingsUrlsState]{
		ID:     authSettingsUrlsID,
		Output: state,
	}, nil
}

func (AuthSettingsUrls) Read(ctx context.Context, req infer.ReadRequest[AuthSettingsUrlsArgs, AuthSettingsUrlsState]) (infer.ReadResponse[AuthSettingsUrlsArgs, AuthSettingsUrlsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[AuthSettingsUrlsArgs, AuthSettingsUrlsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := user_authentication_settings.Get(ctx, service)
	if err != nil {
		return infer.ReadResponse[AuthSettingsUrlsArgs, AuthSettingsUrlsState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[AuthSettingsUrlsArgs, AuthSettingsUrlsState]{}, fmt.Errorf("couldn't read auth settings URLs")
	}

	state := AuthSettingsUrlsState{
		AuthSettingsUrlsArgs: AuthSettingsUrlsArgs{Urls: resp.URLs},
		ResourceId:           authSettingsUrlsID,
	}
	return infer.ReadResponse[AuthSettingsUrlsArgs, AuthSettingsUrlsState]{
		ID:     authSettingsUrlsID,
		Inputs: state.AuthSettingsUrlsArgs,
		State:  state,
	}, nil
}

func (AuthSettingsUrls) Update(ctx context.Context, req infer.UpdateRequest[AuthSettingsUrlsArgs, AuthSettingsUrlsState]) (infer.UpdateResponse[AuthSettingsUrlsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[AuthSettingsUrlsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	urls := user_authentication_settings.ExemptedUrls{URLs: req.Inputs.Urls}
	if _, err := user_authentication_settings.Update(ctx, service, urls); err != nil {
		return infer.UpdateResponse[AuthSettingsUrlsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(2 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[AuthSettingsUrlsState]{}, activationErr
		}
	}

	state := AuthSettingsUrlsState{
		AuthSettingsUrlsArgs: req.Inputs,
		ResourceId:           authSettingsUrlsID,
	}
	return infer.UpdateResponse[AuthSettingsUrlsState]{Output: state}, nil
}

func (AuthSettingsUrls) Delete(ctx context.Context, req infer.DeleteRequest[AuthSettingsUrlsState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (AuthSettingsUrls) Annotate(a infer.Annotator) {
	describeResource(a, &AuthSettingsUrls{}, `The zia_auth_settings_urls resource manages the URLs that are exempted from user authentication in the Zscaler Internet Access (ZIA) cloud service. This singleton resource allows you to define a list of URLs that bypass the ZIA authentication process.

For more information, see the [ZIA User Authentication documentation](https://help.zscaler.com/zia/authentication-exemptions).

{{% examples %}}
## Example Usage

{{% example %}}
### Basic Authentication Settings URLs

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.AuthSettingsUrls("example", {
    urls: [
        ".example.com",
        ".internal.corp.com",
    ],
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.AuthSettingsUrls("example",
    urls=[
        ".example.com",
        ".internal.corp.com",
    ],
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:AuthSettingsUrls
    properties:
      urls:
        - .example.com
        - .internal.corp.com
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

## Import

This is a singleton resource and does not support traditional import. It is automatically managed by the provider.
`)
}

func (a *AuthSettingsUrlsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.Urls, "List of URLs that are exempted from user authentication.")
}

func (s *AuthSettingsUrlsState) Annotate(a infer.Annotator) {
	a.Describe(&s.ResourceId, "The internal resource identifier for the authentication settings URLs.")
}

var _ infer.CustomResource[AuthSettingsUrlsArgs, AuthSettingsUrlsState] = AuthSettingsUrls{}

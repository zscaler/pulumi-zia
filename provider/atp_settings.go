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

// Package provider implements the AtpSettings (Advanced Threat Protection Settings) resource.
// Adopted from terraform-provider-zia resource_zia_atp_settings.go (resourceAdvancedThreatSettings).
// Singleton: UpdateAdvancedThreatSettings for create/update, GetAdvancedThreatSettings for read. Delete no-op.

package provider

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/advancedthreatsettings"
)

const atpSettingsID = "advanced_threat_settings"

// AtpSettings implements the zia:index:AtpSettings resource.
type AtpSettings struct{}

// AtpSettingsArgs are the inputs.
type AtpSettingsArgs struct {
	RiskTolerance                        *int     `pulumi:"riskTolerance,optional"`
	RiskToleranceCapture                 *bool    `pulumi:"riskToleranceCapture,optional"`
	CmdCtlServerBlocked                  *bool    `pulumi:"cmdCtlServerBlocked,optional"`
	CmdCtlServerCapture                  *bool    `pulumi:"cmdCtlServerCapture,optional"`
	CmdCtlTrafficBlocked                  *bool    `pulumi:"cmdCtlTrafficBlocked,optional"`
	CmdCtlTrafficCapture                  *bool    `pulumi:"cmdCtlTrafficCapture,optional"`
	MalwareSitesBlocked                   *bool    `pulumi:"malwareSitesBlocked,optional"`
	MalwareSitesCapture                   *bool    `pulumi:"malwareSitesCapture,optional"`
	ActivexBlocked                       *bool    `pulumi:"activexBlocked,optional"`
	ActivexCapture                       *bool    `pulumi:"activexCapture,optional"`
	BrowserExploitsBlocked               *bool    `pulumi:"browserExploitsBlocked,optional"`
	BrowserExploitsCapture               *bool    `pulumi:"browserExploitsCapture,optional"`
	FileFormatVunerabilitesBlocked       *bool    `pulumi:"fileFormatVunerabilitesBlocked,optional"`
	FileFormatVunerabilitesCapture        *bool    `pulumi:"fileFormatVunerabilitesCapture,optional"`
	KnownPhishingSitesBlocked             *bool    `pulumi:"knownPhishingSitesBlocked,optional"`
	KnownPhishingSitesCapture            *bool    `pulumi:"knownPhishingSitesCapture,optional"`
	SuspectedPhishingSitesBlocked         *bool    `pulumi:"suspectedPhishingSitesBlocked,optional"`
	SuspectedPhishingSitesCapture         *bool    `pulumi:"suspectedPhishingSitesCapture,optional"`
	SuspectAdwareSpywareSitesBlocked      *bool    `pulumi:"suspectAdwareSpywareSitesBlocked,optional"`
	SuspectAdwareSpywareSitesCapture      *bool    `pulumi:"suspectAdwareSpywareSitesCapture,optional"`
	WebSpamBlocked                       *bool    `pulumi:"webSpamBlocked,optional"`
	WebSpamCapture                       *bool    `pulumi:"webSpamCapture,optional"`
	IrcTunnellingBlocked                 *bool    `pulumi:"ircTunnellingBlocked,optional"`
	IrcTunnellingCapture                 *bool    `pulumi:"ircTunnellingCapture,optional"`
	AnonymizerBlocked                    *bool    `pulumi:"anonymizerBlocked,optional"`
	AnonymizerCapture                    *bool    `pulumi:"anonymizerCapture,optional"`
	CookieStealingBlocked                *bool    `pulumi:"cookieStealingBlocked,optional"`
	CookieStealingPcapEnabled            *bool    `pulumi:"cookieStealingPcapEnabled,optional"`
	PotentialMaliciousRequestsBlocked    *bool    `pulumi:"potentialMaliciousRequestsBlocked,optional"`
	PotentialMaliciousRequestsCapture   *bool    `pulumi:"potentialMaliciousRequestsCapture,optional"`
	BlockedCountries                     []string `pulumi:"blockedCountries,optional"`
	BlockCountriesCapture                *bool    `pulumi:"blockCountriesCapture,optional"`
	BitTorrentBlocked                    *bool    `pulumi:"bitTorrentBlocked,optional"`
	BitTorrentCapture                    *bool    `pulumi:"bitTorrentCapture,optional"`
	TorBlocked                           *bool    `pulumi:"torBlocked,optional"`
	TorCapture                           *bool    `pulumi:"torCapture,optional"`
	GoogleTalkBlocked                    *bool    `pulumi:"googleTalkBlocked,optional"`
	GoogleTalkCapture                    *bool    `pulumi:"googleTalkCapture,optional"`
	SshTunnellingBlocked                 *bool    `pulumi:"sshTunnellingBlocked,optional"`
	SshTunnellingCapture                 *bool    `pulumi:"sshTunnellingCapture,optional"`
	CryptoMiningBlocked                  *bool    `pulumi:"cryptoMiningBlocked,optional"`
	CryptoMiningCapture                  *bool    `pulumi:"cryptoMiningCapture,optional"`
	AdSpywareSitesBlocked                *bool    `pulumi:"adSpywareSitesBlocked,optional"`
	AdSpywareSitesCapture                *bool    `pulumi:"adSpywareSitesCapture,optional"`
	DgaDomainsBlocked                    *bool    `pulumi:"dgaDomainsBlocked,optional"`
	DgaDomainsCapture                    *bool    `pulumi:"dgaDomainsCapture,optional"`
	AlertForUnknownSuspiciousC2Traffic   *bool    `pulumi:"alertForUnknownSuspiciousC2Traffic,optional"`
	MaliciousUrlsCapture                 *bool    `pulumi:"maliciousUrlsCapture,optional"`
}

// AtpSettingsState is the persisted state.
type AtpSettingsState struct {
	AtpSettingsArgs
	ResourceId string `pulumi:"resourceId"`
}

func (AtpSettings) Create(ctx context.Context, req infer.CreateRequest[AtpSettingsArgs]) (infer.CreateResponse[AtpSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.CreateResponse[AtpSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := atpSettingsArgsToAPI(req.Inputs)
	if _, _, err := advancedthreatsettings.UpdateAdvancedThreatSettings(ctx, service, apiReq); err != nil {
		return infer.CreateResponse[AtpSettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(1 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.CreateResponse[AtpSettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	resp, err := advancedthreatsettings.GetAdvancedThreatSettings(ctx, service)
	if err != nil {
		state := AtpSettingsState{
			AtpSettingsArgs: req.Inputs,
			ResourceId:      atpSettingsID,
		}
		return infer.CreateResponse[AtpSettingsState]{
			ID:     atpSettingsID,
			Output: state,
		}, nil
	}
	state := atpSettingsAPIToState(resp)
	return infer.CreateResponse[AtpSettingsState]{
		ID:     atpSettingsID,
		Output: state,
	}, nil
}

func (AtpSettings) Read(ctx context.Context, req infer.ReadRequest[AtpSettingsArgs, AtpSettingsState]) (infer.ReadResponse[AtpSettingsArgs, AtpSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.ReadResponse[AtpSettingsArgs, AtpSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	resp, err := advancedthreatsettings.GetAdvancedThreatSettings(ctx, service)
	if err != nil {
		return infer.ReadResponse[AtpSettingsArgs, AtpSettingsState]{}, err
	}
	if resp == nil {
		return infer.ReadResponse[AtpSettingsArgs, AtpSettingsState]{}, fmt.Errorf("couldn't read advanced threat settings")
	}

	state := atpSettingsAPIToState(resp)
	return infer.ReadResponse[AtpSettingsArgs, AtpSettingsState]{
		ID:     atpSettingsID,
		Inputs: state.AtpSettingsArgs,
		State:  state,
	}, nil
}

func (AtpSettings) Update(ctx context.Context, req infer.UpdateRequest[AtpSettingsArgs, AtpSettingsState]) (infer.UpdateResponse[AtpSettingsState], error) {
	cfg := infer.GetConfig[Config](ctx)
	if cfg.Client() == nil {
		return infer.UpdateResponse[AtpSettingsState]{}, fmt.Errorf("ZIA provider not configured")
	}
	service := cfg.Client().Service

	apiReq := atpSettingsArgsToAPI(req.Inputs)
	if _, _, err := advancedthreatsettings.UpdateAdvancedThreatSettings(ctx, service, apiReq); err != nil {
		return infer.UpdateResponse[AtpSettingsState]{}, err
	}

	if shouldActivate() {
		time.Sleep(1 * time.Second)
		if activationErr := triggerActivation(ctx, cfg.Client()); activationErr != nil {
			return infer.UpdateResponse[AtpSettingsState]{}, activationErr
		}
	} else {
		log.Printf("[INFO] Skipping configuration activation due to ZIA_ACTIVATION env var not being set to true.")
	}

	resp, err := advancedthreatsettings.GetAdvancedThreatSettings(ctx, service)
	if err != nil {
		state := AtpSettingsState{
			AtpSettingsArgs: req.Inputs,
			ResourceId:      atpSettingsID,
		}
		return infer.UpdateResponse[AtpSettingsState]{Output: state}, nil
	}
	return infer.UpdateResponse[AtpSettingsState]{Output: atpSettingsAPIToState(resp)}, nil
}

func (AtpSettings) Delete(ctx context.Context, req infer.DeleteRequest[AtpSettingsState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (AtpSettings) Annotate(a infer.Annotator) {
	describeResource(a, &AtpSettings{}, `The zia_atp_settings resource manages Advanced Threat Protection (ATP) settings in the Zscaler Internet Access (ZIA) cloud service. ATP settings control which threat categories are blocked or captured (logged) for packet capture analysis. This is a singleton resource.

For more information, see the [ZIA Advanced Threat Protection documentation](https://help.zscaler.com/zia/about-advanced-threat-protection-policy).

{{% examples %}}
## Example Usage

{{% example %}}
### Configure ATP Settings

`+tripleBacktick("typescript")+`
import * as zia from "@bdzscaler/pulumi-zia";

const example = new zia.AtpSettings("example", {
    malwareSitesBlocked: true,
    malwareSitesCapture: true,
    knownPhishingSitesBlocked: true,
    knownPhishingSitesCapture: true,
    cmdCtlServerBlocked: true,
    cryptoMiningBlocked: true,
    torBlocked: true,
    riskTolerance: 0,
});
`+tripleBacktick("")+`

`+tripleBacktick("python")+`
import zscaler_pulumi_zia as zia

example = zia.AtpSettings("example",
    malware_sites_blocked=True,
    malware_sites_capture=True,
    known_phishing_sites_blocked=True,
    known_phishing_sites_capture=True,
    cmd_ctl_server_blocked=True,
    crypto_mining_blocked=True,
    tor_blocked=True,
    risk_tolerance=0,
)
`+tripleBacktick("")+`

`+tripleBacktick("yaml")+`
resources:
  example:
    type: zia:AtpSettings
    properties:
      malwareSitesBlocked: true
      malwareSitesCapture: true
      knownPhishingSitesBlocked: true
      knownPhishingSitesCapture: true
      cmdCtlServerBlocked: true
      cryptoMiningBlocked: true
      torBlocked: true
      riskTolerance: 0
`+tripleBacktick("")+`

{{% /example %}}
{{% /examples %}}

> This is a singleton resource. Import is not applicable.
`)
}

func (a *AtpSettingsArgs) Annotate(ann infer.Annotator) {
	ann.Describe(&a.RiskTolerance, "Risk tolerance level. Controls the sensitivity for detecting threats.")
	ann.Describe(&a.RiskToleranceCapture, "Enable packet capture for risk tolerance events.")
	ann.Describe(&a.CmdCtlServerBlocked, "Block known command-and-control servers.")
	ann.Describe(&a.CmdCtlServerCapture, "Enable packet capture for command-and-control server events.")
	ann.Describe(&a.CmdCtlTrafficBlocked, "Block command-and-control traffic.")
	ann.Describe(&a.CmdCtlTrafficCapture, "Enable packet capture for command-and-control traffic events.")
	ann.Describe(&a.MalwareSitesBlocked, "Block known malware sites.")
	ann.Describe(&a.MalwareSitesCapture, "Enable packet capture for malware site events.")
	ann.Describe(&a.ActivexBlocked, "Block ActiveX controls.")
	ann.Describe(&a.ActivexCapture, "Enable packet capture for ActiveX events.")
	ann.Describe(&a.BrowserExploitsBlocked, "Block browser exploits.")
	ann.Describe(&a.BrowserExploitsCapture, "Enable packet capture for browser exploit events.")
	ann.Describe(&a.FileFormatVunerabilitesBlocked, "Block file format vulnerabilities.")
	ann.Describe(&a.FileFormatVunerabilitesCapture, "Enable packet capture for file format vulnerability events.")
	ann.Describe(&a.KnownPhishingSitesBlocked, "Block known phishing sites.")
	ann.Describe(&a.KnownPhishingSitesCapture, "Enable packet capture for known phishing site events.")
	ann.Describe(&a.SuspectedPhishingSitesBlocked, "Block suspected phishing sites.")
	ann.Describe(&a.SuspectedPhishingSitesCapture, "Enable packet capture for suspected phishing site events.")
	ann.Describe(&a.SuspectAdwareSpywareSitesBlocked, "Block suspect adware/spyware sites.")
	ann.Describe(&a.SuspectAdwareSpywareSitesCapture, "Enable packet capture for suspect adware/spyware site events.")
	ann.Describe(&a.WebSpamBlocked, "Block web spam.")
	ann.Describe(&a.WebSpamCapture, "Enable packet capture for web spam events.")
	ann.Describe(&a.IrcTunnellingBlocked, "Block IRC tunnelling.")
	ann.Describe(&a.IrcTunnellingCapture, "Enable packet capture for IRC tunnelling events.")
	ann.Describe(&a.AnonymizerBlocked, "Block anonymizers.")
	ann.Describe(&a.AnonymizerCapture, "Enable packet capture for anonymizer events.")
	ann.Describe(&a.CookieStealingBlocked, "Block cookie stealing attempts.")
	ann.Describe(&a.CookieStealingPcapEnabled, "Enable packet capture for cookie stealing events.")
	ann.Describe(&a.PotentialMaliciousRequestsBlocked, "Block potentially malicious requests.")
	ann.Describe(&a.PotentialMaliciousRequestsCapture, "Enable packet capture for potentially malicious request events.")
	ann.Describe(&a.BlockedCountries, "List of countries (ISO 3166-1 alpha-2 codes) to block.")
	ann.Describe(&a.BlockCountriesCapture, "Enable packet capture for blocked countries events.")
	ann.Describe(&a.BitTorrentBlocked, "Block BitTorrent traffic.")
	ann.Describe(&a.BitTorrentCapture, "Enable packet capture for BitTorrent events.")
	ann.Describe(&a.TorBlocked, "Block Tor traffic.")
	ann.Describe(&a.TorCapture, "Enable packet capture for Tor events.")
	ann.Describe(&a.GoogleTalkBlocked, "Block Google Talk traffic.")
	ann.Describe(&a.GoogleTalkCapture, "Enable packet capture for Google Talk events.")
	ann.Describe(&a.SshTunnellingBlocked, "Block SSH tunnelling.")
	ann.Describe(&a.SshTunnellingCapture, "Enable packet capture for SSH tunnelling events.")
	ann.Describe(&a.CryptoMiningBlocked, "Block crypto mining traffic.")
	ann.Describe(&a.CryptoMiningCapture, "Enable packet capture for crypto mining events.")
	ann.Describe(&a.AdSpywareSitesBlocked, "Block adware/spyware sites.")
	ann.Describe(&a.AdSpywareSitesCapture, "Enable packet capture for adware/spyware site events.")
	ann.Describe(&a.DgaDomainsBlocked, "Block domain generation algorithm (DGA) domains.")
	ann.Describe(&a.DgaDomainsCapture, "Enable packet capture for DGA domain events.")
	ann.Describe(&a.AlertForUnknownSuspiciousC2Traffic, "Enable alerts for unknown or suspicious C2 traffic.")
	ann.Describe(&a.MaliciousUrlsCapture, "Enable packet capture for malicious URL events.")
}

func (s *AtpSettingsState) Annotate(a infer.Annotator) {
	a.Describe(&s.ResourceId, "The internal resource identifier for the ATP settings.")
}

var _ infer.CustomResource[AtpSettingsArgs, AtpSettingsState] = AtpSettings{}

func atpSettingsArgsToAPI(in AtpSettingsArgs) advancedthreatsettings.AdvancedThreatSettings {
	use := func(b *bool) bool {
		if b != nil {
			return *b
		}
		return false
	}
	useInt := func(i *int) int {
		if i != nil {
			return *i
		}
		return 0
	}
	processedBlockedCountries := processCountries(in.BlockedCountries)
	return advancedthreatsettings.AdvancedThreatSettings{
		RiskTolerance:                        useInt(in.RiskTolerance),
		RiskToleranceCapture:                 use(in.RiskToleranceCapture),
		CmdCtlServerBlocked:                  use(in.CmdCtlServerBlocked),
		CmdCtlServerCapture:                  use(in.CmdCtlServerCapture),
		CmdCtlTrafficBlocked:                 use(in.CmdCtlTrafficBlocked),
		CmdCtlTrafficCapture:                 use(in.CmdCtlTrafficCapture),
		MalwareSitesBlocked:                  use(in.MalwareSitesBlocked),
		MalwareSitesCapture:                  use(in.MalwareSitesCapture),
		ActiveXBlocked:                       use(in.ActivexBlocked),
		ActiveXCapture:                       use(in.ActivexCapture),
		BrowserExploitsBlocked:               use(in.BrowserExploitsBlocked),
		BrowserExploitsCapture:               use(in.BrowserExploitsCapture),
		FileFormatVulnerabilitiesBlocked:     use(in.FileFormatVunerabilitesBlocked),
		FileFormatVulnerabilitiesCapture:     use(in.FileFormatVunerabilitesCapture),
		KnownPhishingSitesBlocked:            use(in.KnownPhishingSitesBlocked),
		KnownPhishingSitesCapture:            use(in.KnownPhishingSitesCapture),
		SuspectedPhishingSitesBlocked:        use(in.SuspectedPhishingSitesBlocked),
		SuspectedPhishingSitesCapture:        use(in.SuspectedPhishingSitesCapture),
		SuspectAdwareSpywareSitesBlocked:     use(in.SuspectAdwareSpywareSitesBlocked),
		SuspectAdwareSpywareSitesCapture:     use(in.SuspectAdwareSpywareSitesCapture),
		WebspamBlocked:                       use(in.WebSpamBlocked),
		WebspamCapture:                       use(in.WebSpamCapture),
		IrcTunnellingBlocked:                 use(in.IrcTunnellingBlocked),
		IrcTunnellingCapture:                 use(in.IrcTunnellingCapture),
		AnonymizerBlocked:                    use(in.AnonymizerBlocked),
		AnonymizerCapture:                    use(in.AnonymizerCapture),
		CookieStealingBlocked:                use(in.CookieStealingBlocked),
		CookieStealingPCAPEnabled:            use(in.CookieStealingPcapEnabled),
		PotentialMaliciousRequestsBlocked:    use(in.PotentialMaliciousRequestsBlocked),
		PotentialMaliciousRequestsCapture:   use(in.PotentialMaliciousRequestsCapture),
		BlockedCountries:                     processedBlockedCountries,
		BlockCountriesCapture:                use(in.BlockCountriesCapture),
		BitTorrentBlocked:                    use(in.BitTorrentBlocked),
		BitTorrentCapture:                    use(in.BitTorrentCapture),
		TorBlocked:                           use(in.TorBlocked),
		TorCapture:                           use(in.TorCapture),
		GoogleTalkBlocked:                    use(in.GoogleTalkBlocked),
		GoogleTalkCapture:                    use(in.GoogleTalkCapture),
		SshTunnellingBlocked:                 use(in.SshTunnellingBlocked),
		SshTunnellingCapture:                 use(in.SshTunnellingCapture),
		CryptoMiningBlocked:                  use(in.CryptoMiningBlocked),
		CryptoMiningCapture:                  use(in.CryptoMiningCapture),
		AdSpywareSitesBlocked:                use(in.AdSpywareSitesBlocked),
		AdSpywareSitesCapture:                use(in.AdSpywareSitesCapture),
		DgaDomainsBlocked:                    use(in.DgaDomainsBlocked),
		DgaDomainsCapture:                    use(in.DgaDomainsCapture),
		AlertForUnknownOrSuspiciousC2Traffic: use(in.AlertForUnknownSuspiciousC2Traffic),
		MaliciousUrlsCapture:                 use(in.MaliciousUrlsCapture),
	}
}

func atpSettingsAPIToState(r *advancedthreatsettings.AdvancedThreatSettings) AtpSettingsState {
	// Process BlockedCountries: strip "COUNTRY_" prefix for user-facing output
	var processedBlockedCountries []string
	if r.BlockedCountries != nil {
		processedBlockedCountries = make([]string, len(r.BlockedCountries))
		for i, country := range r.BlockedCountries {
			processedBlockedCountries[i] = strings.TrimPrefix(country, "COUNTRY_")
		}
	}
	return AtpSettingsState{
		AtpSettingsArgs: AtpSettingsArgs{
			RiskTolerance:                        intPtr(r.RiskTolerance),
			RiskToleranceCapture:                 boolPtr(r.RiskToleranceCapture),
			CmdCtlServerBlocked:                  boolPtr(r.CmdCtlServerBlocked),
			CmdCtlServerCapture:                  boolPtr(r.CmdCtlServerCapture),
			CmdCtlTrafficBlocked:                 boolPtr(r.CmdCtlTrafficBlocked),
			CmdCtlTrafficCapture:                 boolPtr(r.CmdCtlTrafficCapture),
			MalwareSitesBlocked:                  boolPtr(r.MalwareSitesBlocked),
			MalwareSitesCapture:                  boolPtr(r.MalwareSitesCapture),
			ActivexBlocked:                       boolPtr(r.ActiveXBlocked),
			ActivexCapture:                       boolPtr(r.ActiveXCapture),
			BrowserExploitsBlocked:               boolPtr(r.BrowserExploitsBlocked),
			BrowserExploitsCapture:               boolPtr(r.BrowserExploitsCapture),
			FileFormatVunerabilitesBlocked:       boolPtr(r.FileFormatVulnerabilitiesBlocked),
			FileFormatVunerabilitesCapture:       boolPtr(r.FileFormatVulnerabilitiesCapture),
			KnownPhishingSitesBlocked:            boolPtr(r.KnownPhishingSitesBlocked),
			KnownPhishingSitesCapture:           boolPtr(r.KnownPhishingSitesCapture),
			SuspectedPhishingSitesBlocked:        boolPtr(r.SuspectedPhishingSitesBlocked),
			SuspectedPhishingSitesCapture:        boolPtr(r.SuspectedPhishingSitesCapture),
			SuspectAdwareSpywareSitesBlocked:     boolPtr(r.SuspectAdwareSpywareSitesBlocked),
			SuspectAdwareSpywareSitesCapture:     boolPtr(r.SuspectAdwareSpywareSitesCapture),
			WebSpamBlocked:                       boolPtr(r.WebspamBlocked),
			WebSpamCapture:                       boolPtr(r.WebspamCapture),
			IrcTunnellingBlocked:                 boolPtr(r.IrcTunnellingBlocked),
			IrcTunnellingCapture:                 boolPtr(r.IrcTunnellingCapture),
			AnonymizerBlocked:                    boolPtr(r.AnonymizerBlocked),
			AnonymizerCapture:                    boolPtr(r.AnonymizerCapture),
			CookieStealingBlocked:                boolPtr(r.CookieStealingBlocked),
			CookieStealingPcapEnabled:            boolPtr(r.CookieStealingPCAPEnabled),
			PotentialMaliciousRequestsBlocked:    boolPtr(r.PotentialMaliciousRequestsBlocked),
			PotentialMaliciousRequestsCapture:   boolPtr(r.PotentialMaliciousRequestsCapture),
			BlockedCountries:                     processedBlockedCountries,
			BlockCountriesCapture:               boolPtr(r.BlockCountriesCapture),
			BitTorrentBlocked:                    boolPtr(r.BitTorrentBlocked),
			BitTorrentCapture:                    boolPtr(r.BitTorrentCapture),
			TorBlocked:                           boolPtr(r.TorBlocked),
			TorCapture:                           boolPtr(r.TorCapture),
			GoogleTalkBlocked:                    boolPtr(r.GoogleTalkBlocked),
			GoogleTalkCapture:                    boolPtr(r.GoogleTalkCapture),
			SshTunnellingBlocked:                 boolPtr(r.SshTunnellingBlocked),
			SshTunnellingCapture:                 boolPtr(r.SshTunnellingCapture),
			CryptoMiningBlocked:                  boolPtr(r.CryptoMiningBlocked),
			CryptoMiningCapture:                  boolPtr(r.CryptoMiningCapture),
			AdSpywareSitesBlocked:                boolPtr(r.AdSpywareSitesBlocked),
			AdSpywareSitesCapture:                boolPtr(r.AdSpywareSitesCapture),
			DgaDomainsBlocked:                    boolPtr(r.DgaDomainsBlocked),
			DgaDomainsCapture:                    boolPtr(r.DgaDomainsCapture),
			AlertForUnknownSuspiciousC2Traffic:   boolPtr(r.AlertForUnknownOrSuspiciousC2Traffic),
			MaliciousUrlsCapture:                 boolPtr(r.MaliciousUrlsCapture),
		},
		ResourceId: atpSettingsID,
	}
}

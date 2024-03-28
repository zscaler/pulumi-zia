// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func skipNoZIACreds(t *testing.T) {
	username := os.Getenv("ZIA_USERNAME")
	if username == "" {
		t.Skipf("Skipping test due to missing ZIA_USERNAME variable")
	}
	password := os.Getenv("ZIA_PASSWORD")
	if password == "" {
		t.Skipf("Skipping test due to missing ZIA_PASSWORD variable")
	}
	api_key := os.Getenv("ZIA_API_KEY")
	if api_key == "" {
		t.Skipf("Skipping test due to missing ZIA_API_KEY variable")
	}
	cloud := os.Getenv("ZIA_CLOUD")
	if cloud == "" {
		t.Skipf("Skipping test due to missing ZIA_CLOUD variable")
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
	}
}

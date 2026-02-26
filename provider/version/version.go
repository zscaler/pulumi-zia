package version

// Version is set by the Go linker during build.
var Version string

// PulumiSDKVersion is the fallback Pulumi SDK version used when
// runtime/debug.ReadBuildInfo() is unavailable (e.g. in test binaries).
// Keep in sync with the github.com/pulumi/pulumi/sdk/v3 version in go.mod.
const PulumiSDKVersion = "v3.212.0"

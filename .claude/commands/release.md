---
description: Prepare and verify a new release of the pulumi-zia provider
---

Guide the user through preparing a release of the pulumi-zia native provider. This command should be run from the pulumi-zia repository root.

## Prerequisites Check
1. Verify we're in the pulumi-zia repository
2. Ensure the working tree is clean: `git status`
3. Ensure we're on the `main` branch

## Workflow

### 1. Gather Information
Ask the user:
- What version number? (e.g., "1.0.0", "1.1.0-beta.1")
- Is this a stable release or pre-release?
- Are there any last-minute changes to include?

### 2. Pre-Release Checks

#### Build everything from scratch
```bash
make provider
make generate_schema
make build_sdks
```

#### Run tests
```bash
make test_provider
```

#### Verify schema is up to date
```bash
# Regenerate and check for diffs
make generate_schema
git diff provider/cmd/pulumi-resource-zia/schema.json
```
If there are diffs, the schema was stale — commit the update first.

#### Verify SDK package metadata
Check version placeholders in:
- `sdk/nodejs/package.json` - `version` field
- `sdk/python/pyproject.toml` - `version` field
- `sdk/dotnet/zscaler.PulumiPackage.Zia.csproj` - `<Version>` element

These should show `1.0.0-alpha.0+dev` (replaced at build time by `PROVIDER_VERSION`).

### 3. Verify CI/CD Configuration

#### GoReleaser
```bash
# Dry run to verify config
cd /path/to/repo && goreleaser check
```

Verify `.goreleaser.yml` has:
- `project_name: pulumi-zia`
- `binary: pulumi-resource-zia`
- `main: ./cmd/pulumi-resource-zia/`
- Correct ldflags path: `github.com/zscaler/pulumi-zia/provider/version.Version`

#### GitHub Secrets
Confirm these secrets are configured in the GitHub repository:
- `NPM_TOKEN` - for `@bdzscaler/pulumi-zia`
- `PYPI_API_TOKEN` - for `zscaler_pulumi_zia`
- `NUGET_PUBLISH_KEY` - for `zscaler.PulumiPackage.Zia`

### 4. Create and Push Tag
```bash
git tag v{version}
git push origin v{version}
```

This triggers the `release.yaml` workflow which:
1. Builds provider binaries for darwin/linux/windows (amd64+arm64) via GoReleaser
2. Creates a GitHub Release with the binaries
3. Creates an `sdk/v{version}` tag for Go SDK consumers
4. Builds and publishes SDKs:
   - Node.js to npm
   - Python to PyPI
   - .NET to NuGet
   - Go SDK available via the `sdk/` tag

### 5. Post-Release Verification

#### Check GitHub Actions
- Monitor the release workflow at `https://github.com/zscaler/pulumi-zia/actions`
- Verify all jobs pass (publish_binary, then publish_sdk for each language)

#### Verify published packages
```bash
# npm
npm view @bdzscaler/pulumi-zia version

# PyPI
pip index versions zscaler-pulumi-zia

# NuGet
dotnet package search zscaler.PulumiPackage.Zia --source https://api.nuget.org/v3/index.json
```

#### Verify GitHub Release
- Check `https://github.com/zscaler/pulumi-zia/releases/tag/v{version}`
- Ensure binary assets are attached for all platforms

## Troubleshooting

**GoReleaser fails**: Check `.goreleaser.yml` for boilerplate names — all references must be `pulumi-resource-zia`, not `pulumi-resource-provider-boilerplate`.

**SDK publish fails**: Verify `PROVIDER_VERSION` is set in the release workflow. Without it, SDKs build as `1.0.0-alpha.0+dev`.

**NuGet path error**: NuGet packages are at `sdk/dotnet/bin/Debug/*.nupkg`, not `sdk/dotnet/dotnet/bin/Debug/`.

**Go SDK not resolvable**: Verify the `sdk/v{version}` tag was created and pushed.

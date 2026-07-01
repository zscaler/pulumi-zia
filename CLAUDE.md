# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) and Cursor when working with code in this repository.

## Overview

This is the **Pulumi ZIA (Zscaler Internet Access) native provider**, built with the `pulumi-go-provider` SDK using the `infer` package. Unlike bridged providers that wrap a Terraform provider, this provider is written entirely in Go and communicates directly with the Zscaler API via the `zscaler-sdk-go/v3` SDK. Each resource and data source is a Go struct with CRUD methods implemented by hand.

## Key Architecture Concepts

### Native Provider with `infer`

The provider uses `github.com/pulumi/pulumi-go-provider/infer`, which derives Pulumi resource and function definitions from Go structs via reflection:

- **Resources** are Go structs implementing `infer.CustomCreate`, `infer.CustomRead`, `infer.CustomUpdate`, `infer.CustomDelete`
- **Functions (data sources)** are Go structs implementing `Invoke(ctx, infer.FunctionRequest[Args]) (infer.FunctionResponse[Result], error)`
- **Schema/documentation** is generated automatically from struct tags and `Annotate()` methods
- **Config** implements `infer.CustomConfigure` to initialize the Zscaler SDK client

### Resource Pattern

Every resource follows the same pattern with three structs:

```go
type FooArgs struct {           // Input properties (pulumi tags)
    Name *string `pulumi:"name,optional"`
}

type FooState struct {          // Persisted state (embeds Args + outputs)
    FooArgs
    ResourceId string `pulumi:"resourceId"`
}

type Foo struct{}               // Resource type (implements CRUD + Annotate)
```

- `Args` struct: input properties with `pulumi:"fieldName,optional"` tags
- `State` struct: embeds `Args` and adds output-only fields
- Resource struct: empty struct implementing CRUD methods and `Annotate()`

### Documentation via `Annotate()`

All documentation is embedded in Go code via the `infer.Annotator` interface:

- **Resource descriptions** use `describeResource(a, &Foo{}, "description with Hugo shortcodes")` (value receiver, wraps in `defer recover()` to handle a `pulumi-go-provider` reflection edge case)
- **Property descriptions** use `a.Describe(&a.FieldName, "description")` (pointer receiver)
- **Function descriptions** use `a.Describe(f, "description")` (pointer receiver)
- Multi-language examples use Hugo shortcodes (`{{% examples %}}`, `{{% example %}}`) with the `tripleBacktick()` helper from `utils.go`

### Activation Pattern

ZIA requires explicit policy activation after changes. Resources that modify ZIA policy call `triggerActivation()` after CRUD operations. The `ZIA_ACTIVATION` environment variable controls whether activation is triggered (default: false). The `Activation` resource provides explicit activation control.

### Rule Ordering

Many ZIA rule resources (SSL inspection, URL filtering, firewall filtering/DNS/IPS, NAT, forwarding control, file type control, DLP web, cloud app control, CASB DLP/malware, bandwidth control, sandbox, traffic capture — 15 resources total) have ordering constraints. Ordering is reconciled by a shared, background **diff-based convergent reorder engine** in `common.go` (`reorderAll` / `reorderWithBeforeReorder` / `markOrderRuleAsDone` / `waitForReorder`), plus per-rule-type semaphores and mutexes.

**This engine is a faithful, line-for-line port of the Terraform provider's `reorderAll` (`terraform-provider-zia/zia/common.go`).** It MUST stay in sync with that file. The Terraform implementation is the source of truth — it received the diff-based rewrite in terraform-provider-zia PR #567 ("Significantly reduced apply time and API call volume … across all rule-based resources"), preceded by the race-condition fix in PR #521. When the Terraform reorder logic changes, port the same change here; when adding a new orderable rule resource here, wire it to this engine exactly as the existing ones do.

#### Why a convergence loop is required (the bug fixed by porting this)

The ZIA API treats a rule's `order` as an **insertion/shift**, not an absolute set: PUTting a rule to order 3 shifts every rule currently at ≥3 down by one. Pulumi (like Terraform `--parallel 10`) dispatches resource `Create`s concurrently, so multiple in-flight order writes interleave. A single ascending sweep of absolute-position PUTs therefore does **not** land on the declared ordering — it scrambles custom rules and drifts predefined rules. (See [issue #74](https://github.com/zscaler/pulumi-zia/issues/74): 11 SSL inspection rules at orders 1–11 ended up scrambled, non-deterministic across re-applies. The old, pre-port engine did exactly one reorder pass and never reconciled the residual drift.)

The ported engine fixes this by reconciling to convergence:

- **`getCurrent func() (map[int]OrderRule, error)`** (NOT a count). Each pass calls `getCurrent` **once** to fetch `{ruleID → {Order, Rank}}` for every rule the API knows about. The old engine's `getCount func() (int, error)` callback could not diff and was the root cause — never reintroduce it.
- **Diff-based PUTs.** A rule already at its desired `Order` AND `Rank` is skipped entirely (no GET, no PUT). Only drifted rules are written. A fully-settled pass costs one `GetAll` and zero PUTs.
- **Out-of-range deferral.** Desired orders outside `1..count` (more rules still being POSTed) are deferred to the next tick instead of erroring with `INVALID_INPUT_ARGUMENT`.
- **Convergence requires two consecutive clean passes**, so an in-flight PUT from the previous pass is reflected by `GetAll` before the goroutine returns.
- **Deadlock-breakers** for the parallel-batch case: `maxStuckOnSkippedTicks` (fast early-exit so a blocked `waitForReorder` releases and the next Create batch can extend the orderable range) and `maxNoProgressTicks` (slow safety net for genuinely unreachable declared orders, e.g. gaps).

#### Per-resource wiring (keep identical across all 15 resources)

In each rule resource's `Create` and `Update`, after the initial create/update with a temporary appended order:

1. Call `reorderWithBeforeReorder(OrderRule{Order: intendedOrder, Rank: intendedRank}, id, resourceType, getCurrent, updateOrder, nil)`.
2. `getCurrent` builds `map[int]OrderRule` from the SDK `GetAll`/`GetByRuleType` (filtering out default/predefined rules where applicable, e.g. `filterOutBandwidthDefaultRule`, `filterOutDefaultSandboxRule`). Use `Rank: r.Rank` — except CASB malware rules, which have no admin rank and register `Rank: 0` on both the desired `OrderRule` and the map.
3. `updateOrder` re-fetches the rule, strips server-managed/read-only fields (`LastModifiedTime`, `LastModifiedBy`, `Predefined`, `DefaultRule`, `AccessControl`) so PUTs against predefined rules aren't rejected with "Request body is invalid", sets `Order`/`Rank`, and PUTs.
4. Call `markOrderRuleAsDone(id, resourceType)` then `waitForReorder(resourceType)` before reading back / activating.

### Client Initialization

The Zscaler SDK client is created in `Config.Configure()` and stored in `Config.client`. Resources access it via `infer.GetConfig[Config](ctx).Client()`. The client wrapper lives in `provider/internal/zia/client.go` and supports:

- OAuth2 authentication (client ID + client secret, or client ID + private key)
- Sandbox-only authentication (sandbox token)
- Environment variable fallbacks for all credentials
- HTTP proxy configuration
- Custom user agent string with Pulumi SDK version

## Common Development Commands

### Building

```bash
# Build the provider binary
make provider

# Build individual SDKs
make build_nodejs
make build_python
make build_dotnet
make build_go
make build_java

# Build everything (provider + all SDKs)
make build
```

### Schema Generation

```bash
# Generate schema.json from the provider binary
make generate_schema

# Generate all SDK code from schema
make codegen
```

The schema is generated by running `pulumi package get-schema ./bin/pulumi-resource-zia` and saved to `provider/cmd/pulumi-resource-zia/schema.json`. SDK code is then generated from this schema using `pulumi package gen-sdk`.

### Testing

```bash
# Run provider unit tests
make test_provider

# Run all tests including e2e (requires ZIA credentials)
make test

# Run specific e2e test
cd examples && go test -v -run TestUrlFilteringRule -tags=all -timeout 2h
```

E2E tests require these environment variables:
- `ZSCALER_CLIENT_ID`, `ZSCALER_CLIENT_SECRET`, `ZSCALER_VANITY_DOMAIN`, `ZSCALER_CLOUD`
- Or legacy: `ZIA_USERNAME`, `ZIA_PASSWORD`, `ZIA_API_KEY`, `ZIA_CLOUD`

### SDK Development

```bash
# Generate all SDKs
make codegen

# Link Node.js SDK locally for testing
make install_nodejs_sdk
cd examples/simple
yarn link @bdzscaler/pulumi-zia
pulumi up
```

### Linting

```bash
make lint
```

## Project Structure

- `provider/` - Provider implementation in Go (single Go module at root)
  - `cmd/pulumi-resource-zia/` - Provider binary entrypoint and committed `schema.json`
  - `internal/zia/` - Zscaler SDK client wrapper (`client.go`)
  - `version/` - Version variable set by linker flags
  - `config.go` - Provider configuration and client initialization
  - `provider.go` - Resource and function registration
  - `utils.go` - Shared helpers (`tripleBacktick`, `describeResource`, `triggerActivation`, etc.)
  - `common.go` - Rule ordering logic, expand/flatten helpers
  - `validator.go` - Input validation functions
  - `*.go` - One file per resource or data source (100 files total)
- `sdk/` - Generated SDK code
  - `nodejs/` - TypeScript/JavaScript SDK (`@bdzscaler/pulumi-zia`)
  - `python/` - Python SDK (`zscaler_pulumi_zia`)
  - `dotnet/` - .NET SDK (`zscaler.PulumiPackage.Zia`)
  - `go/pulumi-zia/` - Go SDK (`github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia`)
  - `java/` - Java SDK
- `tests/` - E2E integration tests (45 test files)
- `examples/` - Example Pulumi programs (simple, yaml, Go, dotnet)
- `local_dev/` - Local development helpers
- `.github/workflows/` - CI/CD (ci.yaml, release.yaml)
- `.goreleaser.yml` - GoReleaser configuration for binary releases

## Version Management

- `PROVIDER_VERSION` - Defaults to `1.0.0-alpha.0+dev` locally; set from git tags in CI
- `VERSION_GENERIC` - Normalized version derived via `pulumictl convert-version`
- Version is embedded in the binary via `-ldflags "-X github.com/zscaler/pulumi-zia/provider/version.Version=..."`
- SDK versions are set during `pulumi package gen-sdk --version`

## Zscaler Go SDK (Source of Truth)

All resources and data sources are backed by `github.com/zscaler/zscaler-sdk-go/v3`. The SDK is vendored at `vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/`. **Any new resource or attribute must exist in the SDK first** -- the provider wraps SDK structs and CRUD functions, it does not call the Zscaler API directly.

When adding or modifying a resource, always start by examining the SDK:
```bash
# Find the SDK package for a service
ls vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/

# Read the SDK struct and CRUD functions
cat vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/<package>/<file>.go
```

The SDK defines:
- **Structs** with `json` tags -- these determine the available fields and types
- **CRUD functions** (Create, Get, GetByName, Update, Delete, GetAll) -- these determine the provider's CRUD method signatures
- **Validation** -- some constraints come from the SDK, others from `validator.go`

When the SDK is updated (new fields, new resources), update the vendor directory first:
```bash
go get github.com/zscaler/zscaler-sdk-go/v3@latest
go mod vendor
```

The Terraform provider (`terraform-provider-zia`) can also serve as a secondary reference for attribute descriptions, validation logic, and behavioral nuances, but the SDK is the authoritative source for types and API surface.

## Adding New Resources

1. **Check the SDK**: Find the struct and CRUD functions in `vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/<package>/`
2. Create a new file `provider/<resource_name>.go` with `Args`, `State`, and resource structs that map to the SDK struct fields
3. Implement `Create`, `Read`, `Update`, `Delete` methods calling the SDK CRUD functions
4. Add `Annotate()` methods on the resource (value receiver using `describeResource`), Args (pointer receiver), and State (pointer receiver) with descriptions and Hugo shortcode examples
5. Register with `infer.Resource(ResourceName{})` in `provider/provider.go` `WithResources()`
6. Rebuild and regenerate: `make provider && make generate_schema`

## Adding New Data Sources (Functions)

1. **Check the SDK**: Find the Get/GetByName/GetAll functions and return types in `vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/<package>/`
2. Create a new file `provider/<datasource_name>.go` with `Args`, `Result`, and function structs that map to the SDK return type fields
3. Implement `Invoke(ctx, infer.FunctionRequest[Args]) (infer.FunctionResponse[Result], error)` calling the SDK getter functions
4. Add `Annotate()` methods (all pointer receivers) with descriptions
5. Register with `infer.Function(&FunctionName{})` in `provider/provider.go` `WithFunctions()`
6. Rebuild and regenerate: `make provider && make generate_schema`

## CI/CD

- **CI** (`ci.yaml`): Builds provider, generates schema, runs unit tests, builds all SDKs on push/PR
- **Release** (`release.yaml`): Triggered by `v*.*.*` tags; builds binaries via GoReleaser, publishes SDKs to npm/PyPI/NuGet, creates `sdk/` tag for Go SDK
- Provider binaries are distributed via GitHub Releases (not S3)

## Important Implementation Notes

### `describeResource` Helper

Resource-level `Annotate()` uses value receivers (required by `pulumi-go-provider`), but the `GetToken` code path creates non-addressable `reflect.Value`s causing panics on `a.Describe(&T{}, ...)`. The `describeResource()` helper in `utils.go` wraps the call in `defer recover()` so the schema generation path still captures descriptions correctly.

### Struct Tag Conventions

- `pulumi:"fieldName"` - Required field
- `pulumi:"fieldName,optional"` - Optional field
- `provider:"secret"` - Marks field as secret (used for credentials in Config)

### SDK Package Names

- npm: `@bdzscaler/pulumi-zia`
- PyPI: `zscaler_pulumi_zia`
- NuGet: `zscaler.PulumiPackage.Zia`
- Go: `github.com/zscaler/pulumi-zia/sdk/go/pulumi-zia`

### License

MIT License - Copyright (c) 2023 Zscaler Technology Alliances

---
description: Add a new ZIA resource to the native Pulumi provider
---

Guide the user through adding a new ZIA resource to the pulumi-zia native provider. This command should be run from the pulumi-zia repository root.

## Prerequisites Check
1. Verify we're in the pulumi-zia repository
2. Check that `bin/pulumi-resource-zia` exists or can be built

## Workflow

### 1. Gather Information
Ask the user:
- What is the resource name? (e.g., "FirewallFilteringRule", "DlpDictionary")
- Which Zscaler SDK service package does it use? List available packages for the user:
  ```bash
  ls vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/
  ```
- Is there an equivalent Terraform resource to reference? (optional, for descriptions/validation)
- Does this resource require activation after CRUD operations?
- Does this resource have rule ordering requirements?

### 2. Check SDK Types (Source of Truth)

**The `zscaler-sdk-go/v3` SDK is the authoritative source for all resource fields and CRUD functions.** New resources and attributes must exist in the SDK before they can be added to the provider.

- Read the SDK struct and CRUD functions:
  ```bash
  cat vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/<package>/*.go
  ```
- Identify all struct fields (the `json` tags define field names, Go types define Pulumi types)
- Identify which fields are required vs optional (pointers = optional, value types = required)
- Identify the Create/Get/GetByName/Update/Delete function signatures
- If the SDK doesn't have the needed resource/field, tell the user they need to update the SDK first:
  ```bash
  go get github.com/zscaler/zscaler-sdk-go/v3@latest
  go mod vendor
  ```

The Terraform provider can be used as a secondary reference for attribute descriptions and behavioral details, but the SDK defines the API surface.

### 3. Create the Resource File
Create `provider/<resource_name>.go` with:

#### MIT License Header
```go
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
```

#### Resource Structs
- `Args` struct with all input properties (use `pulumi:"fieldName,optional"` for optional fields)
- `State` struct embedding `Args` plus output-only fields (e.g., `ResourceId`)
- Empty resource struct

#### CRUD Methods
- `Create(ctx, name, input, preview)` - Create and return initial state
- `Read(ctx, id, inputs, state)` - Read current state from API
- `Update(ctx, id, olds, news, preview)` - Update and return new state
- `Delete(ctx, id, state)` - Delete the resource

#### Annotate Methods
- **Resource** (value receiver): Use `describeResource(a, &Resource{}, "...")` with Hugo shortcode examples in TypeScript, Python, Go, YAML
- **Args** (pointer receiver): Describe every input property
- **State** (pointer receiver): Describe output-only properties

#### Pattern for client access:
```go
cfg := infer.GetConfig[Config](ctx)
if cfg.Client() == nil {
    return state, fmt.Errorf("ZIA provider not configured")
}
svc := cfg.Client().Service
```

### 4. Register the Resource
- Add `infer.Resource(ResourceName{})` to the `WithResources()` list in `provider/provider.go`
- Keep the list organized (alphabetical or by category)

### 5. Build and Verify
```bash
# Compile
go build ./provider/...

# Build binary
make provider

# Generate schema
make generate_schema

# Verify resource appears in schema
jq '.resources["zia:index:ResourceName"]' provider/cmd/pulumi-resource-zia/schema.json | head -5
```

### 6. Verify Documentation
- Check that the resource description includes Hugo shortcodes
- Verify property descriptions are in the schema:
```bash
jq '.resources["zia:index:ResourceName"].inputProperties | keys' provider/cmd/pulumi-resource-zia/schema.json
```

## Common Patterns

**Activation**: If the resource modifies ZIA policy:
```go
if shouldActivate() {
    if err := triggerActivation(ctx, cfg.Client()); err != nil {
        log.Printf("[ERROR] Activation failed: %s", err)
    }
}
```

**Rule ordering**: If the resource has an `Order` field, check `common.go` for the appropriate lock and reorder helper.

**Retry on error**: Use `failFastOnErrorCodes(err)` to detect non-retryable errors.

## Reference
- Existing resource example: `provider/url_filtering_rules.go`
- Utils: `provider/utils.go` (`describeResource`, `tripleBacktick`, `triggerActivation`)
- Common: `provider/common.go` (rule ordering, expand/flatten helpers)
- CLAUDE.md for full project context

---
description: Add a new ZIA data source (function/invoke) to the native Pulumi provider
---

Guide the user through adding a new ZIA data source to the pulumi-zia native provider. This command should be run from the pulumi-zia repository root.

## Prerequisites Check
1. Verify we're in the pulumi-zia repository
2. Check that `bin/pulumi-resource-zia` exists or can be built

## Workflow

### 1. Gather Information
Ask the user:
- What is the data source name? (e.g., "GetLocationGroup", "GetDatacenters")
- Which Zscaler SDK service package does it use? List available packages:
  ```bash
  ls vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/
  ```
- Is there an equivalent Terraform data source to reference? (optional, for descriptions)
- Does it look up by ID, name, or other fields? Or does it return a list?

### 2. Check SDK Types (Source of Truth)

**The `zscaler-sdk-go/v3` SDK is the authoritative source.** Data sources wrap SDK getter functions and return SDK structs.

- Read the SDK structs and getter functions:
  ```bash
  cat vendor/github.com/zscaler/zscaler-sdk-go/v3/zscaler/zia/services/<package>/*.go
  ```
- Identify the return type struct fields (the `json` tags define field names)
- Identify available lookup methods (Get, GetByID, GetByName, GetAll, etc.)
- If the SDK doesn't have the needed getter function, tell the user they need to update the SDK first:
  ```bash
  go get github.com/zscaler/zscaler-sdk-go/v3@latest
  go mod vendor
  ```

### 3. Create the Data Source File
Create `provider/<datasource_name>.go` with:

#### MIT License Header
(Same as resource - see `add-resource` command)

#### Structs
- `GetFooArgs` - Input/lookup fields (use `*type` with `pulumi:"field,optional"` for optional lookups)
- `GetFooResult` - All output fields
- `GetFoo` - Empty function struct

#### Invoke Method
```go
func (*GetFoo) Invoke(ctx context.Context, req infer.FunctionRequest[GetFooArgs]) (infer.FunctionResponse[GetFooResult], error) {
    cfg := infer.GetConfig[Config](ctx)
    if cfg.Client() == nil {
        return infer.FunctionResponse[GetFooResult]{}, fmt.Errorf("ZIA provider not configured")
    }
    svc := cfg.Client().Service
    // Lookup logic...
}
```

#### Annotate Methods (ALL pointer receivers)
- Function: `a.Describe(f, "Use this data source to...")`
- Args: describe each lookup field
- Result: describe each output field

### 4. Register the Function
- Add `infer.Function(&GetFoo{})` to `WithFunctions()` in `provider/provider.go`
- **Important**: Pass a pointer (`&GetFoo{}`), not a value

### 5. Build and Verify
```bash
go build ./provider/...
make provider
make generate_schema

# Verify function appears in schema
jq '.functions["zia:index:getFoo"]' provider/cmd/pulumi-resource-zia/schema.json | head -5
```

## Key Differences from Resources
- Use `infer.Function(&GetFoo{})` (pointer) not `infer.Resource(Foo{})`  (value)
- ALL Annotate methods use pointer receivers
- No CRUD methods — only `Invoke`
- Result struct instead of State struct
- No need for activation or ordering

## Lookup Patterns

**Single resource by ID or name:**
```go
var resp *sdk.Foo
if req.Input.Id != nil && *req.Input.Id != 0 {
    r, err := sdk.GetFoo(ctx, svc, *req.Input.Id)
    if err != nil { return ..., err }
    resp = r
}
if resp == nil && req.Input.Name != nil && *req.Input.Name != "" {
    r, err := sdk.GetFooByName(ctx, svc, *req.Input.Name)
    if err != nil { return ..., err }
    resp = r
}
```

**List with optional filtering (e.g., GetDatacenters):**
```go
all, err := sdk.GetAll(ctx, svc)
// Apply filters from req.Input
// Return filtered list
```

## Reference
- Existing single-lookup example: `provider/location_groups.go`
- Existing list example: `provider/datacenters.go`
- CLAUDE.md for full project context

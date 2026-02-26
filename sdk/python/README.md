# Pulumi Native Provider for Zscaler Internet Access (ZIA)

A **native** Pulumi provider for Zscaler Internet Access (ZIA), built on [pulumi-go-provider](https://github.com/pulumi/pulumi-go-provider) and the official [zscaler-sdk-go](https://github.com/zscaler/zscaler-sdk-go). This provider does **not** use the Terraform bridge.

## Requirements

- Go 1.24+
- Pulumi CLI
- ZIA credentials (OAuth2 or legacy)

## Build & Install

```bash
# Build the provider
make provider

# Generate SDKs (Go, Node.js, Python, Dotnet, Java)
make codegen

# Install provider binary to $GOPATH/bin
make install
```

Note: Dotnet SDK build may fail due to namespace collision; Go, Node.js, and Python SDKs are verified.

## Python SDK

```bash
# Install the Python SDK locally (from repo root)
pip install -e sdk/python
```

Or add to your project's `requirements.txt`:
```
pulumi>=3.165.0,<4.0.0
-e /path/to/pulumi-zia-native/sdk/python
```

## Provider Configuration

Configure via Pulumi config or environment variables:

| Config | Env Var | Description |
|--------|---------|-------------|
| `clientId` | `ZSCALER_CLIENT_ID` | OAuth2 client ID |
| `clientSecret` | `ZSCALER_CLIENT_SECRET` | OAuth2 client secret (secret) |
| `vanityDomain` | `ZSCALER_VANITY_DOMAIN` | Zscaler vanity domain |
| `cloud` | `ZSCALER_CLOUD` | Zscaler cloud (optional) |
| `debug` | `ZSCALER_SDK_LOG` + `ZSCALER_SDK_VERBOSE` | Enable SDK API request/response logging |

Or use `clientId` + `privateKey` + `vanityDomain` for key-based auth.

## Debugging & Troubleshooting

To see Zscaler API calls and provider internals:

**1. Provider config (recommended)** – set `debug: true` in your Pulumi config:
```yaml
config:
  zia:debug: true
```

**2. Environment variables** – same behavior as Terraform provider:
```bash
export ZSCALER_SDK_LOG=true
export ZSCALER_SDK_VERBOSE=true
pulumi up
```

**2b. Log to file** – Pulumi often does not forward provider stdout. To capture SDK logs to a file:
```bash
export ZSCALER_SDK_LOG=true
export ZSCALER_SDK_VERBOSE=true
export ZSCALER_SDK_LOG_FILE=/tmp/zia-sdk.log
pulumi up --yes
```
Then inspect the log: `cat /tmp/zia-sdk.log`

**3. Pulumi verbose logging** – provider output may be hidden unless Pulumi forwards it. Use:
```bash
pulumi up --logtostderr -v=9
```
For maximum visibility, capture both stdout and stderr:
```bash
pulumi up --logtostderr -v=9 2>&1 | tee debug.log
```

**4. Pulumi provider debugging** – attach a debugger to the provider:
```bash
PULUMI_DEBUG_PROVIDERS="zia:12345" pulumi up
```
Then attach your Go debugger (e.g. Delve) to port 12345.

## Resources

| Resource | Pulumi Token | Description |
|----------|--------------|-------------|
| RuleLabel | `zia:index:RuleLabel` | ZIA rule label for organizing firewall/URL filtering rules |

## Example

**Node.js (TypeScript):**
```typescript
import { RuleLabel } from "@zia/zia/provider";

const label = new RuleLabel("my-label", {
    name: "pulumi-managed-label",
    description: "Created by Pulumi",
});

export const ruleLabelId = label.ruleLabelId;
```

**Python:**
```python
import pulumi
from zia_zia import RuleLabel

label = RuleLabel("my-label",
    name="pulumi-managed-label",
    description="Created by Pulumi",
)

pulumi.export("rule_label_id", label.rule_label_id)
```

## Import

```bash
# Import by ID
pulumi import zia:index:RuleLabel my-label 12345

# Import by name
pulumi import zia:index:RuleLabel my-label "My Label Name"
```

## Documentation

- [mapping.md](docs/mapping.md) - Terraform to Pulumi resource mapping
- [sdk-usage.md](docs/sdk-usage.md) - How each resource uses zscaler-sdk-go
- [progress.md](docs/progress.md) - Implementation progress and next steps

## License

Apache 2.0

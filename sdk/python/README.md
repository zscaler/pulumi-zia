# Zscaler Internet Access (ZIA) Resource Provider

The ZIA Resource Provider lets you manage [ZIA](http://github.com/zscaler/pulumi-zia) resources. To use
this package, please [install the Pulumi CLI first](https://pulumi.com/).

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @bdzscaler/pulumi-zia
```

or `yarn`:

```bash
yarn add @bdzscaler/pulumi-zia
```

### Python

To use from Python, install using `pip`:

```bash
pip install zscaler-pulumi-zia
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/zscaler/pulumi-zia/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package zscaler.PulumiPackage.Zia
```

## Configuration

The following configuration points are available for the `zia` provider:

- `zia:username` (client id: `ZIA_USERNAME`) - (Required) This is the API username to interact with the ZIA cloud.
- `zia:password` (client secret: `ZIA_PASSWORD`) - (Required) This is the password for the API username to authenticate in the ZIA cloud.
- `zia:api_key` (customer id: `ZIA_API_KEY`) - (Required) This is the API Key used in combination with the ``username`` and ``password``
- `zia:zia_cloud` (cloud environment: `ZIA_CLOUD`) - (Required) The cloud name where the ZIA tenant is hosted. The supported values are:
  - ``zscaler``
  - ``zscalerone``
  - ``zscalertwo``
  - ``zscalerthree``
  - ``zscloud``
  - ``zscalerbeta``
  - ``zscalergov``

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/zia/api-docs/).

## Support

This template/solution are released under an as-is, best effort, support
policy. These scripts should be seen as community supported and Zscaler
Business Development Team will contribute our expertise as and when possible.
We do not provide technical support or help in using or troubleshooting the components
of the project through our normal support options such as Zscaler support teams,
or ASC (Authorized Support Centers) partners and backline
support options. The underlying product used (Zscaler Internet Access API) by the
scripts or templates are still supported, but the support is only for the
product functionality and not for help in deploying or using the template or
script itself. Unless explicitly tagged, all projects or work posted in our
GitHub repository at (<https://github.com/zscaler>) or sites other
than our official Downloads page on <https://support.zscaler.com>
are provided under the best effort policy.

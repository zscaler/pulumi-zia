"""A Python Pulumi program"""

import zscaler_pulumi_zia as zia

# Pre-Shared-Key is "FAKE" used for testing only.
vpnCredentials = zia.ZIATrafficForwardingVPNCredentials("vpn-credentials-example",
    comments = "Pulumi VPN Credentials",
    type = "UFQDN",
    pre_shared_key = "<YOUR_PRESHARED_KEY_HERE>",
    fqdn = "sjc-100@securitygeek.io",
)

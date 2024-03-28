"""A Python Pulumi program"""

import zscaler_pulumi_zia as zia

sourceIPGroup = zia.ZIAFirewallFilteringSourceGroups("source-ip-group-example",
    name = "Pulumi IP Source Group",
    description = "Pulumi IP Source Group",
    ip_addresses = ["192.168.1.1", "192.168.1.2", "192.168.1.3"],
)

"""A Python Pulumi program"""

import zscaler_pulumi_zia as zia

destFQDNGroup = zia.ZIAFirewallFilteringDestinationGroups("destination-ip-group-example",
    name = "Pulumi IP Destination Group",
    description = "Pulumi IP Destination Group",
    type = "DSTN_FQDN",
    addresses = [ "test1.acme.com", "test2.acme.com", "test3.acme.com" ],
)

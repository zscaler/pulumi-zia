"""A Python Pulumi program"""

import zscaler_pulumi_zia as zia

staticIP = zia.ZIATrafficForwardingStaticIP("static-ip-example",
    comment = "Pulumi Static IP",
    ip_address = "121.234.54.105",
    geo_override = True,
	latitude = -36.848461,
    longitude = 174.763336,
    routable_ip = True,
)
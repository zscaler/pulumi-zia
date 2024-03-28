"""A Python Pulumi program"""

import zscaler_pulumi_zia as zia

staticIP = zia.ZIARuleLabels("rule-labels-example",
    name = "Pulumi Rule Label",
    description = "Pulumi Rule Label",
)
"""ZIA RuleLabel example using the native Pulumi ZIA provider Python SDK."""
import pulumi
from zscaler_pulumi_zia import RuleLabel

# Create a ZIA Rule Label
label = RuleLabel(
    "my-label",
    name="pulumi-managed-label",
    description="Created by Pulumi native ZIA provider (Python)",
)

pulumi.export("rule_label_id", label.rule_label_id)
pulumi.export("name", label.name)

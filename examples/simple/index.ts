import * as pulumi from "@pulumi/pulumi";
import { RuleLabel } from "@bdzscaler/pulumi-zia";

// Create a ZIA Rule Label
const label = new RuleLabel("my-label", {
    name: "pulumi-managed-label",
    description: "Created by Pulumi native ZIA provider",
});

export const ruleLabelId = label.ruleLabelId;
export const name = label.name;

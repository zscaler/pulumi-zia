import * as zia from "@zscaler/pulumi-zia";

// ZIA Traffic Forwarding - Static IP
const staticIP = new zia.trafficforwarding.TrafficForwardingStaticIP("static_ip_example", {
    comment: "Pulumi Traffic Forwarding Static IP",
    geoOverride: true,
    ipAddress: "123.234.244.245",
    latitude: -36.848461,
    longitude: 174.763336,
    routableIp: true,
});
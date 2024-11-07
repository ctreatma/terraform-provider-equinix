---
subcategory: "Fabric"
---

# equinix_fabric_connection_route_filter (Resource)

Fabric V4 API compatible resource allows attachment of Route Filter Polices to Fabric Connections

Additional Documentation:
* Getting Started: https://docs.equinix.com/en-us/Content/Interconnection/FCR/FCR-route-filters.htm
* API: https://developer.equinix.com/dev-docs/fabric/api-reference/fabric-v4-apis#route-filters

## Example Usage

```terraform
resource "equinix_fabric_connection_route_filter" "policy_attachment" {
  connection_id = "<connection_uuid>"
  route_filter_id = "<route_filter_policy_uuid>"
  direction = "INBOUND"
}

output "connection_route_filter_id" {
  value = equinix_fabric_connection_route_filter.policy_attachment.id
}

output "connection_route_filter_connection_id" {
  value = equinix_fabric_connection_route_filter.policy_attachment.connection_id
}

output "connection_route_filter_direction" {
  value = equinix_fabric_connection_route_filter.policy_attachment.direction
}

output "connection_route_filter_type" {
  value = equinix_fabric_connection_route_filter.policy_attachment.type
}

output "connection_route_filter_attachment_status" {
  value = equinix_fabric_connection_route_filter.policy_attachment.attachment_status
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_id` (String) Equinix Assigned UUID of the Equinix Connection to attach the Route Filter Policy to
- `direction` (String) Direction of the filtering of the attached Route Filter Policy
- `route_filter_id` (String) Equinix Assigned UUID of the Route Filter Policy to attach to the Equinix Connection

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `attachment_status` (String) Status of the Route Filter Policy attachment lifecycle
- `href` (String) URI to the attached Route Filter Policy on the Connection
- `id` (String) The ID of this resource.
- `type` (String) Route Filter Type. One of [ "BGP_IPv4_PREFIX_FILTER", "BGP_IPv6_PREFIX_FILTER" ]
- `uuid` (String) Equinix Assigned ID for Route Filter Policy

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
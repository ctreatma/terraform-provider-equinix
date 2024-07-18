---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "equinix_fabric_port Data Source - terraform-provider-equinix"
subcategory: "Fabric"
description: |-
  Fabric V4 API compatible data resource that allow user to fetch port by uuid
---

# equinix_fabric_port (Data Source)

Fabric V4 API compatible data resource that allow user to fetch port by uuid

Additional documentation:
* Getting Started: https://docs.equinix.com/en-us/Content/Interconnection/Fabric/IMPLEMENTATION/fabric-ports-implement.htm
* API: https://developer.equinix.com/dev-docs/fabric/api-reference/fabric-v4-apis#ports

## Example Usage

```terraform
data "equinix_fabric_port" "port_data_name" {
  uuid = "<uuid_of_port>"
}

output "id" {
  value = data.equinix_fabric_port.port_data_name.id
}

output "name" {
  value = data.equinix_fabric_port.port_data_name.name
}

output "state" {
  value = data.equinix_fabric_port.port_data_name.state
}

output "account_name" {
  value = data.equinix_fabric_port.port_data_name.account.0.account_name
}

output "type" {
  value = data.equinix_fabric_port.port_data_name.type
}

output "bandwidth" {
  value = data.equinix_fabric_port.port_data_name.bandwidth
}

output "used_bandwidth" {
  value = data.equinix_fabric_port.port_data_name.used_bandwidth
}

output "encapsulation_type" {
  value = data.equinix_fabric_port.port_data_name.encapsulation.0.type
}

output "ibx" {
  value = data.equinix_fabric_port.port_data_name.location.0.ibx
}

output "metro_code" {
  value = data.equinix_fabric_port.port_data_name.location.0.metro_code
}

output "metro_name" {
  value = data.equinix_fabric_port.port_data_name.location.0.metro_name
}

output "region" {
  value = data.equinix_fabric_port.port_data_name.location.0.region
}

output "device_redundancy_enabled" {
  value = data.equinix_fabric_port.port_data_name.device.0.redundancy.0.enabled
}

output "device_redundancy_priority" {
  value = data.equinix_fabric_port.port_data_name.device.0.redundancy.0.priority
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `uuid` (String) Equinix-assigned port identifier

### Read-Only

- `account` (Set of Object) Customer account information that is associated with this port (see [below for nested schema](#nestedatt--account))
- `available_bandwidth` (Number) Port available bandwidth in Mbps
- `bandwidth` (Number) Port bandwidth in Mbps
- `change_log` (Set of Object) Captures port lifecycle change information (see [below for nested schema](#nestedatt--change_log))
- `description` (String) Port description
- `device` (Set of Object) Port device (see [below for nested schema](#nestedatt--device))
- `encapsulation` (Set of Object) Port encapsulation protocol (see [below for nested schema](#nestedatt--encapsulation))
- `href` (String) Port URI information
- `id` (String) The ID of this resource.
- `lag_enabled` (Boolean) Port Lag
- `location` (Set of Object) Port location information (see [below for nested schema](#nestedatt--location))
- `name` (String) Port name
- `operation` (Set of Object) Port specific operational data (see [below for nested schema](#nestedatt--operation))
- `redundancy` (Set of Object) Port redundancy information (see [below for nested schema](#nestedatt--redundancy))
- `service_type` (String) Port service type
- `state` (String) Port state
- `type` (String) Port type
- `used_bandwidth` (Number) Port used bandwidth in Mbps

<a id="nestedatt--account"></a>

### Nested Schema for `account`

Read-Only:

- `account_name` (String)
- `account_number` (Number)
- `global_cust_id` (String)
- `global_org_id` (String)
- `global_organization_name` (String)
- `org_id` (Number)
- `organization_name` (String)

<a id="nestedatt--change_log"></a>

### Nested Schema for `change_log`

Read-Only:

- `created_by` (String)
- `created_by_email` (String)
- `created_by_full_name` (String)
- `created_date_time` (String)
- `deleted_by` (String)
- `deleted_by_email` (String)
- `deleted_by_full_name` (String)
- `deleted_date_time` (String)
- `updated_by` (String)
- `updated_by_email` (String)
- `updated_by_full_name` (String)
- `updated_date_time` (String)

<a id="nestedatt--device"></a>

### Nested Schema for `device`

Read-Only:

- `name` (String)
- `redundancy` (Set of Object) (see [below for nested schema](#nestedobjatt--device--redundancy))

<a id="nestedobjatt--device--redundancy"></a>

### Nested Schema for `device.redundancy`

Read-Only:

- `enabled` (Boolean)
- `group` (String)
- `priority` (String)

<a id="nestedatt--encapsulation"></a>

### Nested Schema for `encapsulation`

Read-Only:

- `tag_protocol_id` (String)
- `type` (String)

<a id="nestedatt--location"></a>

### Nested Schema for `location`

Read-Only:

- `ibx` (String)
- `metro_code` (String)
- `metro_name` (String)
- `region` (String)

<a id="nestedatt--operation"></a>

### Nested Schema for `operation`

Read-Only:

- `connection_count` (Number)
- `op_status_changed_at` (String)
- `operational_status` (String)

<a id="nestedatt--redundancy"></a>

### Nested Schema for `redundancy`

Read-Only:

- `enabled` (Boolean)
- `group` (String)
- `priority` (String)
---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_ha_mgmt_update"
description: |-
  Configure fortiadc HA mgmt parameters configuration.
---

# fortiadc_system_ha
Configure fortiadc HA mgmt parameters configuration.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `mgmt-interface` - Management interface. 
* `mgmt-ip-allowaccess` - Allow access with the management ip address. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .
* `mgmt-ip` - IP address of management interface. 
* `mgmt-status` - Management status enable/disable. Valid values: enable/disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.


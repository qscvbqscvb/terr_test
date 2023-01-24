---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_access_list_child_rule"
description: |-
  Configure fortiadc Configure ipv4 access list..
---

# fortiadc_router_access_list_child_rule
Configure fortiadc Configure ipv4 access list..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - id..
* `action` - Action.. Valid values: 1:permit, 2:deny .
* `prefix` - Prefix to define regular filter criteria.. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Access List Child Rule can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_access_list_child_rule.labelname {{mkey}}
```

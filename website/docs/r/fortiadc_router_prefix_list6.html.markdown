---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_prefix_list6"
description: |-
  Configure fortiadc Configure ipv6 prefix list..
---

# fortiadc_router_prefix_list6
Configure fortiadc Configure ipv6 prefix list..

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - name of ipv6 prefix list..
* `description` - Comments.. 



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Prefix List6 can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_prefix_list6.labelname {{mkey}}
```

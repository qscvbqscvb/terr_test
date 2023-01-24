---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_policy"
description: |-
  Configure fortiadc route policy.
---

# fortiadc_router_policy
Configure fortiadc route policy.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The number of route policy.
* `destination` - Destination for this route. 
* `source` - Source for this route policy. 
* `gateway` - Gateway for this route. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Router Policy can be imported using any of these accepted formats:
```
$ terraform import fortiadc_router_policy.labelname {{mkey}}
```

---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_isp_addr"
description: |-
  Configure fortiadc isp address group.
---

# fortiadc_system_isp_addr
Configure fortiadc isp address group.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the addr isp.
* `type` - Type of isp-addr. Valid values: 1:predef, 2:restored, 4:userdef .





## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 System Isp Addr can be imported using any of these accepted formats:
```
$ terraform import fortiadc_system_isp_addr.labelname {{mkey}}
```

---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_compression_child_uri_list"
description: |-
  Configure fortiadc load-balance compression info.
---

# fortiadc_load_balance_compression_child_uri_list
Configure fortiadc load-balance compression info.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - uri list id.
* `uri` - uri string. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import
 Load Balance Compression Child Uri List can be imported using any of these accepted formats:
```
$ terraform import fortiadc_load_balance_compression_child_uri_list.labelname {{mkey}}
```

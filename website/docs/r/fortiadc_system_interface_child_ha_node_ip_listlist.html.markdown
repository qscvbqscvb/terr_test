---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface_child_ha_node_ip_list"
description: |-
  Provides a list of fortiadc system interface child ha node ip list
---

# Data Source: fortiadc_system_interface_child_ha_node_ip_list
Provides a list of `fortiadc system interface child ha node ip list`.

## Example Usage

```hcl
 data "fortiadc_system_interface_child_ha_node_ip_listlist" sample1 {
  filter = "mkey=1"
}

output output1 {
  value = data.fortiadc_system_interface_child_ha_node_ip_listlist.sample1.mkey_list
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](TODO: append doc filter page)

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `mkey_list` -  A list of the `fortiadc_system_interface_child_ha_node_ip_list`.

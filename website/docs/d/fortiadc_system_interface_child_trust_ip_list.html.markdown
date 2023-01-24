---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface_child_trust_ip_list"
description: |-
  Get information on an fortiadc system interface child trust ip list
---

# Data Source: fortiadc_system_interface_child_trust_ip_list
Use this data source to get information on an fortiadc system interface child trust ip list

## Example Usage

```hcl
 data "fortiadc_system_interface_child_trust_ip_list" sample1 {
  name = "1"
}

output output1 {
  value = data.fortiadc_system_interface_child_trust_ip_list.sample1
}
```

## Argument Reference
* `name` - (Required) Specify the mkey of the desired  system interface child trust ip list.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - trust ip name.
* `ip6-max` - end trust IP6. 
* `ip6-min` - start trust IP6. 
* `ip-max` - end trust ip. 
* `ip-netmask` - trust ip/netmask. 
* `ip6-netmask` - trust ip6/netmask. 
* `type` - trust ip type. 
* `ip-min` - start trust IP. 


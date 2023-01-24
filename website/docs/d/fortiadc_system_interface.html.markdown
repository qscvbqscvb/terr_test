---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_interface"
description: |-
  Get information on an fortiadc system interface
---

# Data Source: fortiadc_system_interface
Use this data source to get information on an fortiadc system interface

## Example Usage

```hcl
 data "fortiadc_system_interface" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_interface.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system interface.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - interface name.
* `status` - interface status. 
* `aggregate-mode` - aggregate mode. 
* `dhcp_gw_override` - enable/disable to override gateway. 
* `secondary-ip` - enable/disable to set secondary ip. 
* `floating-ip` - floating ip address of interface. 
* `floating` - enable/disable to set floating ip. 
* `speed` - speed. 
* `redundant-member` - redundant/aggregate interface slaves. 
* `dhcp_ip_overlap` - dhcp_ip_is_overlap. 
* `dhcp_gw_distance` - distance for dhcp_gateway. (1,255)
* `wccp` - enable/disable WCCP. 
* `default-gw` - retrieve default gateway from pppoe server. 
* `trust-ip` - enable/disable to set trust ip. 

* `vlanid` - Vlan ID. (1,4094)


* `vdom` - vdom name. 
* `allowaccess` - Allow access with the interface. 
* `dns-server-override` - use pppoe server DNS or not. 
* `type` - interface type. 
* `ip` - ip address of interface. 
* `traffic-group` - traffic group name. 
* `disc-retry-timeout` - pppoe discovery retry timeout. (1,255)
* `dhcp_gateway` - gateway. 
* `ip6` - ipv6 address of interface. 
* `dedicate-to-management` - dedicate to managemnet interface. 
* `interface` - physical interface name. 
* `password` - pppoe account password. 



* `mtu` - maximum transportation unit. (68,9000)
* `username` - pppoe account user name. 
* `floating-ip6` - floating ipv6 address of interface. 
* `mode` - static,pppoe or dhcp. 
* `aggregate-algorithm` - aggregate interface slave selection algorithm. 


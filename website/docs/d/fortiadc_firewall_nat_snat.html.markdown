---
subcategory: "FortiADC Firewall"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_firewall_nat_snat"
description: |-
  Get information on an fortiadc firewall nat snat
---

# Data Source: fortiadc_firewall_nat_snat
Use this data source to get information on an fortiadc firewall nat snat

## Example Usage

```hcl
 data "fortiadc_firewall_nat_snat" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_firewall_nat_snat.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  firewall nat snat.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - The name of the snat.
* `status` - (en|dis)able snat status. 
* `out-interface` - Output interface name. 
* `from` - Source network. 
* `traffic-group` - traffic group name. 
* `trans-to-ip-end` - Trans-to-pool ip max. 
* `to` - Destination network. 
* `trans-to-type` - Trans to type ip or ip pool or no-nat. 
* `trans-to-ip-start` - Trans-to-pool ip min. 
* `trans-to-ip` - Trans-to ip. 


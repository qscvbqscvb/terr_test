---
subcategory: "FortiADC Router"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_router_bgp"
description: |-
  Get information on an fortiadc router bgp
---

# Data Source: fortiadc_router_bgp
Use this data source to get information on an fortiadc router bgp

## Example Usage

```hcl
 data "fortiadc_router_bgp" sample1 {
}

output output1 {
  value = data.fortiadc_router_bgp.sample1
}
```

## Argument Reference
* `router-id` - (Required) Specify the mkey of the desired  router bgp.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `redistribute-ospf` - Enable/Disable redistribute ospf. 
* `redistribute-connected` - Enable/Disable redistribute connected. 

* `deterministic-med` - Enable/disable enforce deterministic comparison of MED.. 


* `redistribute-connected6` - Enable/Disable redistribute connected6. 
* `as` - Router AS number.. 


* `redistribute-static6` - Enable/Disable redistribute static6.. 
* `always-compare-med` - Enable/disable always compare MED.. 
* `bestpath-cmp-routerid` - Enable/disable compare router ID for identical EBGP paths.. 
* `redistribute-static` - Enable/Disable redistribute static.. 



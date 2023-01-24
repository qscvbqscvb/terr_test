---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_profile_child_client_request_header_erase"
description: |-
  Get information on an fortiadc load balance profile child client request header erase
---

# Data Source: fortiadc_load_balance_profile_child_client_request_header_erase
Use this data source to get information on an fortiadc load balance profile child client request header erase

## Example Usage

```hcl
 data "fortiadc_load_balance_profile_child_client_request_header_erase" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_profile_child_client_request_header_erase.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance profile child client request header erase.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - header erase id.
* `type` - header erase type. 
* `string` - field identify which header to be erased. 


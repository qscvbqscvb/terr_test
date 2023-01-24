---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_pagespeed_child_page_control"
description: |-
  Get information on an fortiadc load balance pagespeed child page control
---

# Data Source: fortiadc_load_balance_pagespeed_child_page_control
Use this data source to get information on an fortiadc load balance pagespeed child page control

## Example Usage

```hcl
 data "fortiadc_load_balance_pagespeed_child_page_control" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_pagespeed_child_page_control.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance pagespeed child page control.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - ID.
* `type` - include/exclue type. 
* `uri` - full uri wildcards, such as http(s)://*example.com/*/htmls/*.html. 


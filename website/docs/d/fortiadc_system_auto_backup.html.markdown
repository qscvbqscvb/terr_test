---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_auto_backup"
description: |-
  Get information on an fortiadc system auto backup
---

# Data Source: fortiadc_system_auto_backup
Use this data source to get information on an fortiadc system auto backup

## Example Usage

```hcl
 data "fortiadc_system_auto_backup" sample1 {
}

output output1 {
  value = data.fortiadc_system_auto_backup.sample1
}
```

## Argument Reference
* `storage` - (Required) Specify the mkey of the desired  system auto backup.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `username` - username for local server. 
* `overwrite-config` - overwrite previous configuration when backup space is full. 
* `config-password` - password for config. 
* `scheduled-backup-status` - enable/disable auto backup. 
* `scheduled-backup-frequency` - scheduled backup frequency. 
* `scheduled-backup-time` - hour and minute, hh: 0-23, mm: {00|15|30|45}. 
* `address` - local server ip address. 
* `scheduled-backup-day` - day of the week: Sunday to Saturday. 
* `folder` - local directory. 
* `password` - password for local server. 
* `port` - local server port. (1,65535)


---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_sdn_connector"
description: |-
  Get information on an fortiadc system sdn connector
---

# Data Source: fortiadc_system_sdn_connector
Use this data source to get information on an fortiadc system sdn connector

## Example Usage

```hcl
 data "fortiadc_system_sdn_connector" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_sdn_connector.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system sdn connector.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - sdn connector name.
* `status` - enable/disable sdn connector. 
* `oci-region-type` - sdn connector oci region type. 
* `port` - server port. (1,65535)
* `oci-cert` - sdn connector oci cert. 
* `sap-ms-http-port` - sap message server http port. (1,65535)
* `aws_secretkey` - sdn connector secret token. 
* `update_interval` - interval (seconds). (30,3600)
* `type` - sdn connector type. 
* `sap-icm-http-port` - sap icm http port. (1,65535)
* `secret_token` - sdn connector secret token. 
* `oci-tenant-id` - sdn connector oci tenant id. 
* `oci-region` - sdn connector oci region. 
* `aws_region` - sdn connector aws region. 
* `oci-user-id` - sdn connector oci user id. 
* `aws_accesskey` - sdn connector secret token. 
* `use-metadata-iam` - enable/disable using iam role. 
* `sap-dns-suffix` - dns name for sap system. 
* `oci-ha-status` - enable/disable use HA service. 
* `sap-password` - sap soap sid admin password. 
* `sap-sidadm` - sap soap sid admin. 
* `oci-compartment-id` - sdn connector oci compartment id. 
* `server` - server ip/FQDN/host(host is only for sap). 


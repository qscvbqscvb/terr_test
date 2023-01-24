---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_global"
description: |-
  Get information on an fortiadc system global
---

# Data Source: fortiadc_system_global
Use this data source to get information on an fortiadc system global

## Example Usage

```hcl
 data "fortiadc_system_global" sample1 {
}

output output1 {
  value = data.fortiadc_system_global.sample1
}
```

## Argument Reference
* `https-server-cert` - (Required) Specify the mkey of the desired  system global.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `vdom-admin` - enable/disable . 
* `telnet-port` - the port number of the telnet service. (1,65535)
* `pre-login-banner` - enable/disable pre-login-banner. 
* `hostname` - appliance's host name. 
* `config-sync-enable` - config sync enable. 


* `vdom-mode` - set vdom mode. 


* `sys-global-language` - Global GUI display language. 
* `gui-log` - Enable/disable menu log on Web GUI. 
* `ssh-port` - the port number of the ssh service. (1,65535)
* `config-sync-port` - config sync port. (1,65535)

* `gui-device-latitude` - latitude between (-90, 90). 
* `gui-router` - Enable/disable menu router on Web GUI. 
* `default-intermediate-ca-group` - appliance's default ssl certificate chain. 
* `admin-idle-timeout` - the idle time-out for system administration. (1,480)
* `https-port` - the port number of the https service. (1,65535)
* `https-redirect` - HTTPS redirect enable/disable. 
* `gui-system` - Enable/disable menu system on Web GUI. 
* `http-port` - the port number of the http service. (1,65535)
* `use-default-hostname` - use-default-hostname enable/disable. 

* `gui-device-longtitude` - longitude between (-180, 180). 


---
subcategory: "FortiADC Load"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_load_balance_virtual_server"
description: |-
  Get information on an fortiadc load balance virtual server
---

# Data Source: fortiadc_load_balance_virtual_server
Use this data source to get information on an fortiadc load balance virtual server

## Example Usage

```hcl
 data "fortiadc_load_balance_virtual_server" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_load_balance_virtual_server.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  load balance virtual server.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - virtual-server name.
* `scripting_list` - virtual server scripting list. 
* `status` - enable/maintain/disable virtual server. 
* `packet-fwd-method` - packet forwarding method. 
* `warmrate` - virtual server warm up rate. (1,86400)
* `http2https` - enable/disable redirect HTTP request to HTTPS. 
* `pool` - pool name. 
* `error-page` - error-page name. 
* `fortiview` - enable/disable fortiview. 
* `ztna_profile` - ZTNA profile name. 
* `auth_policy` - authentication policy name. 
* `adfs-published-service` - AD FS published service. 
* `content-rewriting` - content rewriting. 
* `http2https-port` - HTTP service port list for redirecting HTTP to HTTPS. 
* `domain-name` - virtual server domainname on one click gslb server. 
* `protocol` - virtual server protocol numbers. 
* `error-msg` - error message. 
* `port` - virtual server service port. 
* `scripting_flag` - enable/disable virtual server scripting. 
* `comments` - virtual server comments. 
* `content-routing-list` - content routing list. 
* `wccp` - enable/disable redirect HTTP/HTTPS request to WCCP client. 
* `public-ip-type` - address type. 
* `one-click-gslb-server` - enable/disable setting one click gslb server. 
* `pagespeed` - virtual server pagespeed. 
* `connection-limit` - connection-limit. (0,100000000)
* `host-name` - virtual server hostname on one click gslb server. 
* `ssl-mirror-intf` - interface list to mirror. 
* `source-pool-list` - ip pool name. 
* `type` - virtual-server service type. 
* `method` - method name. 
* `persistence` - persistence name. 
* `dos_profile` - DoS profile name. 
* `profile` - profile name. 
* `content-rewriting-list` - content rewriting list. 
* `l2-exception-list` - layer2 exception list. 
* `ips_profile` - IPS profile name. 
* `traffic-group` - traffic group name. 
* `stream_scripting_list` - virtual server scripting list. 
* `address6` - ipv6 address of virtual server. 
* `clone-pool` - clone pool name. 
* `use-azure-lb-backend-ip` - use azure lb backend ip. 
* `public-ip` - public ip address of virtual server. 
* `clone-traffic-type` - the traffic type to be cloned. 
* `address` - ip address of virtual server. 
* `alone` - enable/disable alone mode. 
* `stream_scripting_flag` - enable/disable virtual server stream scripting. 
* `client_ssl_profile` - client SSL profile (needed for https tcps smtp ftp). 
* `schedule-pool-list` - schedule pool name. 
* `ssl-mirror` - enable/disable SSL mirror. 
* `schedule-list` - enable/disable schedule list. 
* `warmup` - virtual server warm up time. (0,86400)
* `addr-type` - address type. 
* `connection-pool` - connection pool name. 
* `azure-lb-backend` - azure lb backend name. 
* `connection-rate-limit` - virtual server connection rate limit(0 - disable). (0,86400)
* `trans-rate-limit` - virtual server transactions rate limit. (0,1048567)
* `traffic-log` - enable/disable traffic log. 
* `av-profile` - antivirus profile name. 
* `public-ip6` - public ip address of virtual server. 
* `interface` - interface name. 
* `captcha-profile` - Captcha profile. 
* `waf-profile` - web application firewall profile name. 
* `content-routing` - content routing. 


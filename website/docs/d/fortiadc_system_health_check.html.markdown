---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_system_health_check"
description: |-
  Get information on an fortiadc system health check
---

# Data Source: fortiadc_system_health_check
Use this data source to get information on an fortiadc system health check

## Example Usage

```hcl
 data "fortiadc_system_health_check" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiadc_system_health_check.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  system health check.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - health-check name.
* `counter-value` - snmp counter value: range(0, 2147483647). 
* `service_name` - the service name for connecting oracle db. 
* `baseDN` - baseDN. 
* `pwd_type` - RADIUS password type. 
* `dest_addr6` - dest-addr6. 
* `mssql-row` - row. 
* `disk` - disk threshold for snmp. (1,99)
* `row` - the row for sql searching. 
* `http_version` - http version. 
* `passive` - FTP passive enable/disable. 
* `mssql-column` - column. 
* `status_code` - status-code. 
* `folder` - Mail folder name. 
* `addr_type` - address type. 
* `mssql-receive-string` - mssql receive string. 
* `oracle-receive-string` - the receive-string for connecting oracle db. 
* `host_addr` - host-addr. 
* `password` - password. 
* `sip_request_type` - type of SIP request. 
* `remote-password` - remote-password. 
* `timeout` - timeout. (1,3600)
* `acct-appid` - diameter acct application id. 

* `port` - port. (0,65535)
* `allow-ssl-version` - SSL/TLS versions allowed. 
* `remote_port` - remote-port. (0,65535)
* `connect-string` - the connect data for connecting oracle db. 
* `domain_name` - domain-name. 
* `version` - snmp version. 
* `product-name` - product name. 
* `nas-ip` - RADIUS NAS IP address. 
* `bindDN` - bindDN. 
* `mem` - mem threshold for snmp. (1,99)
* `host_addr6` - host-addr6. 
* `http_extra_string` - additional-string. 
* `column` - the column for sql seaching. 
* `http_connect` - http connect type. 
* `auth-appid` - diameter auth application id. 
* `agent-type` - snmp agent type. 
* `filter` - filter. 
* `remote_host` - remote-host. 
* `connect-type` - the type of oracle connect data. 
* `method_type` - method type. 
* `ssl-ciphers` - ssl ciphers including Perfect Forward Secrecy. 
* `community` - snmp community. 
* `script` - health-check script. 
* `dest_addr` - dest-addr. 
* `type` - health-check type. 
* `username` - username. 
* `rtsp-method-type` - type of RTSP method. 
* `mssql-send-string` - mssql send string. 
* `string-value` - snmp string value. 
* `down_retry` - retry. (1,10)
* `origin-realm` - origin realm. 
* `origin-host` - origin host. 
* `secret_key` - RADIUS Secret Key. 
* `compare-type` - snmp compare type. 
* `match_type` - match-type. 
* `attribute` - attribute. 
* `local-cert` - Local certificate reference. 
* `mysql-server-type` - type of MySQL server. 
* `file` - File path on FTP server. 
* `dest_addr_type` - destination address type. 
* `cpu-weight` - snmp cpu weight. (0,200)
* `host-ip-addr` - host ip addr. 
* `host-ip6-addr` - host ip6 addr. 
* `hostname` - hostname. 
* `mem-weight` - snmp mem weight. (0,200)
* `sid` - the sid for connecting oracle db. 

* `disk-weight` - snmp disk weight. (0,200)
* `receive_string` - receive-string. 
* `oid` - snmp oid. 
* `value-type` - snmp value type. 
* `radius-reject` - enbale radius reject. 
* `rtsp-describe-url` - URL for DESCRIBE method. 
* `database` - datebase name of the mssql server. 
* `oracle-send-string` - the send string for connecting oracle db. 
* `send_string` - send-string. 
* `interval` - interval. (1,3600)
* `up_retry` - up-retry. (1,10)
* `remote-username` - remote-username. 
* `cpu` - cpu threshold for snmp. (1,99)
* `vendor-id` - vendor id. 


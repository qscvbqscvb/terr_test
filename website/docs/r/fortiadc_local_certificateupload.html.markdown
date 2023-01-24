---
subcategory: "FortiADC System"
layout: "fortiadc"
page_title: "FortiADC: fortiadc_local_certificateupload"
description: |-
  Upload fortiadc local certificate.
---

# fortiadc_local_certificateupload
Upload fortiadc local certificate.

## Example Usage
```hcl
TODO
```

## Argument Reference

The following arguments are supported:

* `mkey` - name.
* `type` - LocalCert/PKCS12/CertKey.
* `passwd` - certificate password.
* `cert` - cert file path.
* `key` - key file path.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.


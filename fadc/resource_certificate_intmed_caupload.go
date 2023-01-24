// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCertificateIntmedCaUpload() *schema.Resource {
	return &schema.Resource{
		Read:   resourceCertificateIntmedCaUploadRead,
		Update: resourceCertificateIntmedCaUploadUpdate,
		Create: resourceCertificateIntmedCaUploadCreate,
		Delete: resourceCertificateIntmedCaUploadDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceCertificateIntmedCaUploadCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing CertificateIntmedCaUpload: type error")
	}

	obj, err := getObjectCertificateIntmedCaUpload(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CertificateIntmedCaUpload resource while getting object: %v", err)
	}

	_, err = c.CreateCertificateIntmedCaUpload(obj, vdom)

	if err != nil {
		return fmt.Errorf("Error creating CertificateIntmedCaUpload resource: %v", err)
	}

	d.SetId(mkey)

	//return resourceCertificateIntmedCaUploadRead(d, m)
	return nil
}
func resourceCertificateIntmedCaUploadUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCertificateIntmedCaUploadRead(d, m)
}
func resourceCertificateIntmedCaUploadDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCertificateIntmedCaUpload(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateIntmedCaUpload resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceCertificateIntmedCaUploadRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadCertificateIntmedCaUpload(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CertificateIntmedCaUpload resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateIntmedCaUpload(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CertificateIntmedCaUpload resource from API: %v", err)
	}
	return nil
}

func flattenCertificateIntmedCaUploadMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateIntmedCaUpload(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	/*if err = d.Set("subject", flattenCertificateIntmedCaUploadSubject(o["subject"], d, "subject", sv)); err != nil {
		if !fortiAPIPatch(o["subject"]) {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("validfrom", flattenCertificateIntmedCaUploadValidfrom(o["validfrom"], d, "validfrom", sv)); err != nil {
		if !fortiAPIPatch(o["validfrom"]) {
			return fmt.Errorf("Error reading validfrom: %v", err)
		}
	}

	if err = d.Set("validto", flattenCertificateIntmedCaUploadValidto(o["validto"], d, "validto", sv)); err != nil {
		if !fortiAPIPatch(o["validto"]) {
			return fmt.Errorf("Error reading validto: %v", err)
		}
	}

	if err = d.Set("type", flattenCertificateIntmedCaUploadType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ca_type", flattenCertificateIntmedCaUploadiCAtype(o["ca_type"], d, "ca_type", sv)); err != nil {
		if !fortiAPIPatch(o["ca_type"]) {
			return fmt.Errorf("Error reading ca_type: %v", err)
		}
	}

	if err = d.Set("issuer", flattenCertificateIntmedCaUploadIssuer(o["issuer"], d, "issuer", sv)); err != nil {
		if !fortiAPIPatch(o["issuer"]) {
			return fmt.Errorf("Error reading issuer: %v", err)
		}
	}

	if err = d.Set("fingerprint", flattenCertificateIntmedCaUploadFingerprint(o["fingerprint"], d, "fingerprint", sv)); err != nil {
		if !fortiAPIPatch(o["fingerprint"]) {
			return fmt.Errorf("Error reading fingerprint: %v", err)
		}
	}*/

	if err = d.Set("mkey", flattenCertificateIntmedCaUploadMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandCertificateIntmedCaUploadCert(d *schema.ResourceData, v interface{}, pre string, sv string, writer *multipart.Writer) (interface{}, error) {
	return v, nil
}

func expandCertificateIntmedCaUploadMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateIntmedCaUploadVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateIntmedCaUpload(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	body := &bytes.Buffer{}
	obj := make(map[string]interface{})
	writer := multipart.NewWriter(body)

	if v, ok := d.GetOk("cert"); ok {
		t, err := expandCertificateIntmedCaUploadCert(d, v, "cert", sv, writer)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "cert", t.(string), true)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandCertificateIntmedCaUploadMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "mkey", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandCertificateIntmedCaUploadVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "vdom", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	writer.Close()
	obj["content_type"] = writer.FormDataContentType()
	obj["data"] = body.Bytes()

	return &obj, nil
}

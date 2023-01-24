package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	//"github.com/terraform-providers/terraform-provider-fortios/fortios"
	"github.com/terraform-providers/terraform-provider-fortiadc/fadc"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		//ProviderFunc: fortios.Provider})
		ProviderFunc: fortiadc.Provider})
}

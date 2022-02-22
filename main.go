package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-wordle/wordle"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: wordle.Provider,
	})
}

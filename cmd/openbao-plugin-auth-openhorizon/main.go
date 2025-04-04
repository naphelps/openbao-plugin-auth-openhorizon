package main

import (
	"log"
	"os"

	openhorizon "github.com/open-horizon/openbao-plugin-auth-openhorizon"
	"github.com/openbao/openbao/api/v2"
	"github.com/openbao/openbao/sdk/v2/plugin"
)

// This plugin provides authentication support for openhorizon users within bao.
//
// It uses OpenBao's framework to interact with the plugin system.
//
// This plugin must be configured by a bao admin through the /config API. Without the config, the plugin
// is unable to function properly.

func main() {
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err := plugin.ServeMultiplex(&plugin.ServeOpts{
		BackendFactoryFunc: openhorizon.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})

	if err != nil {
		log.Fatal(err)
	}
}

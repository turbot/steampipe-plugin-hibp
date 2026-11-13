package main

import (
	"github.com/turbot/steampipe-plugin-hibp/hibp"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hibp.Plugin})
}

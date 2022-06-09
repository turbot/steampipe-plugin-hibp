package hibp

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type HibpConfig struct {
	ApiKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &HibpConfig{}
}

func GetConfig(connection *plugin.Connection) HibpConfig {
	if connection == nil || connection.Config == nil {
		return HibpConfig{}
	}

	config, _ := connection.Config.(HibpConfig)
	return config
}

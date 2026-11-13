package hibp

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type HibpConfig struct {
	ApiKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &HibpConfig{}
}

func GetConfig(connection *plugin.Connection) HibpConfig {
	if connection == nil || connection.GetConfig() == nil {
		return HibpConfig{}
	}

	config, _ := connection.GetConfig().(HibpConfig)
	return config
}

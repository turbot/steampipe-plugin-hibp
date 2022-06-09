package hibp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-hibp",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"hibp_breach":           tableHIBPBreach(),
			"hibp_breached_account": tableHIBPBreachedAccount(),
			"hibp_password":         tableHIBPPassword(),
			"hibp_paste":            tableHIBPPaste(),
		},
	}

	return p
}

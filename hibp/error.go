package hibp

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func ignore404Error(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
	return strings.Contains(err.Error(), "404 Not Found")
}

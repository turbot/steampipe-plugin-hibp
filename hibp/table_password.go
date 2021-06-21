package hibp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
	"gitlab.com/wedtm/go-hibp"
)

func tablePassword() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_password",
		Description: "Password (hashes) tracked by HIBP",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"hash"}),
			Hydrate:    getPasswords,
		},
		Columns: []*plugin.Column{
			{Name: "hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password.", Transform: transform.FromQual("hash")},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

func getPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		panic(err)
	}

	quals := d.KeyColumnQuals
	queryHash := quals["hash"].GetStringValue()

	pwMatch, _, err := client.Passwords.GetExactPasswordBySHA1(queryHash)

	if err != nil {
		panic(err)
	}
	return pwMatch, nil
}

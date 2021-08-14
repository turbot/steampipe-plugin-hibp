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
		List: &plugin.ListConfig{
			KeyColumns: plugin.AnyColumn([]string{"prefix", "hash"}),
			Hydrate:    listPasswords,
		},
		Columns: []*plugin.Column{
			{Name: "prefix", Type: proto.ColumnType_STRING, Description: "The first five characters of the hash", Transform: transform.FromQual("prefix")},
			{Name: "hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

func listPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals

	prefix := quals["hash"].GetStringValue()
	if prefix == "" {
		prefix = quals["prefix"].GetStringValue()
	}

	passwords, _, err := client.Passwords.GetPasswordsBySHA1Prefix(prefix)

	if err != nil {
		return nil, err
	}

	for _, pw := range passwords {
		d.StreamListItem(ctx, pw)
	}
	return nil, nil
}

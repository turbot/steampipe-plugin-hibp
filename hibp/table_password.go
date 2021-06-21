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
			KeyColumns: plugin.AnyColumn([]string{"hash"}),
			Hydrate:    listPasswords,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"hash"}),
			Hydrate:    getPasswords,
		},
		Columns: []*plugin.Column{
			{Name: "hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password.", Hydrate: getPasswords, Transform: transform.From(stitchHash)},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

func stitchHash(ctx context.Context, tf *transform.TransformData) (interface{}, error) {
	pws := tf.HydrateItem.(*hibp.PasswordMatch)
	quals := tf.KeyColumnQuals
	queryHash := quals["hash"].(string)
	prefix := queryHash[:5]

	newPw := &hibp.PasswordMatch{
		Hash:  prefix + pws.Hash,
		Count: pws.Count,
	}

	return newPw, nil
}

func getPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		panic(err)
	}

	quals := d.KeyColumnQuals
	queryHash := quals["hash"].GetStringValue()
	prefix := queryHash[:5]

	allPw, _, err := client.Passwords.GetPasswordsBySHA1(prefix)
	plugin.Logger(ctx).Warn("getPasswords", "allPw", allPw)
	if err != nil {
		panic(err)
	}

	for _, pw := range allPw {
		if prefix+pw.Hash == queryHash {
			return pw, nil
		}
	}

	return nil, nil
}

func listPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		panic(err)
	}

	quals := d.KeyColumnQuals
	queryHash := quals["hash"].GetStringValue()
	prefix := queryHash[:5]

	passwords, _, err := client.Passwords.GetPasswordsBySHA1(prefix)

	if err != nil {
		panic(err)
	}

	for _, password := range passwords {
		d.StreamListItem(ctx, password)
	}
	return nil, nil
}

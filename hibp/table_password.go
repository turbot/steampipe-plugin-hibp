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
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("hash"),
			Hydrate:    getPassword,
		},
		Columns: []*plugin.Column{
			{Name: "prefix", Type: proto.ColumnType_STRING, Description: "The first five characters of the hash", Transform: transform.From(prefixValue)},
			{Name: "hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password.", Transform: transform.From(hashValue)},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

//// HYDRATE FUNCTIONS

func listPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)
	if err != nil {
		plugin.Logger(ctx).Error("hibp_password.listPasswords", "client.error", err)
		return nil, err
	}

	prefix := d.KeyColumnQuals["hash"].GetStringValue()
	if prefix == "" {
		prefix = d.KeyColumnQuals["prefix"].GetStringValue()
	}

	passwords, _, err := client.Passwords.GetPasswordsBySHA1Prefix(prefix)
	if err != nil {
		plugin.Logger(ctx).Error("hibp_password.listPasswords", "api.error", err)
		return nil, err
	}

	for _, pw := range passwords {
		d.StreamListItem(ctx, pw)
	}
	return nil, nil
}

func getPassword(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)
	if err != nil {
		plugin.Logger(ctx).Error("hibp_password.getPassword", "client.error", err)
		return nil, err
	}

	hash := d.KeyColumnQuals["hash"].GetStringValue()
	pwMatch, _, err := client.Passwords.GetExactPasswordBySHA1(hash)
	if err != nil {
		plugin.Logger(ctx).Error("hibp_password.getPassword", "api.error", err)
		return nil, err
	}

	return pwMatch, nil
}

//// TRANSFORM FUNCTIONS

func hashValue(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	pw := d.HydrateItem.(*hibp.PasswordMatch)
	var hash string
	hashQualValue := d.KeyColumnQuals["hash"]
	if hashQualValue != nil {
		hash = hashQualValue.(string)
		return hash, nil
	}

	return pw.Hash, nil
}

func prefixValue(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	pw := d.HydrateItem.(*hibp.PasswordMatch)
	var prefix string
	hashQualValue := d.KeyColumnQuals["prefix"]
	if hashQualValue != nil {
		prefix = hashQualValue.(string)
		return prefix, nil
	}

	prefix = pw.Hash[:5]
	return prefix, nil
}

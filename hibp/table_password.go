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
			{Name: "prefix", Type: proto.ColumnType_STRING, Description: "The first five characters of the hash", Transform: transform.From(transformPrefix), Hydrate: hydratePrefix},
			{Name: "hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password.", Transform: transform.From(transformHash), Hydrate: hydrateHash},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

func transformHash(ctx context.Context, t *transform.TransformData) (interface{}, error) {
	pw := t.HydrateItem.(string)
	plugin.Logger(ctx).Warn("transformHash", "quals", t.KeyColumnQuals, "pw", pw)
	if val, ok := t.KeyColumnQuals["prefix"]; ok {
		plugin.Logger(ctx).Warn("transformHash", "qual", val)
		return val.(string) + pw, nil
	}

	if val, ok := t.KeyColumnQuals["hash"]; ok {
		plugin.Logger(ctx).Warn("transformHash", "qual", val)
		return val.(string), nil
	}
	return nil, nil
}

func transformPrefix(ctx context.Context, t *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Warn("transformPrefix", "quals", t.KeyColumnQuals)
	if val, ok := t.KeyColumnQuals["prefix"]; ok {
		plugin.Logger(ctx).Warn("transformPrefix", "qual", val)
		return val.(string), nil
	}

	if val, ok := t.KeyColumnQuals["hash"]; ok {
		plugin.Logger(ctx).Warn("transformPrefix", "qual", val)
		return val.(string)[:5], nil
	}
	return nil, nil
}

func hydratePrefix(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	value := d.KeyColumnQuals["hash"].GetStringValue()
	if value == "" {
		value = d.KeyColumnQuals["prefix"].GetStringValue()
	}
	plugin.Logger(ctx).Warn("hydratePrefix", "value", value)
	return value[:5], nil
}

func hydrateHash(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	pw := h.Item.(*hibp.PasswordMatch)
	quals := d.KeyColumnQuals
	queryHash := pw.Hash
	if queryHash == "" {
		queryHash = quals["prefix"].GetStringValue()
	}
	result := queryHash[:5] + pw.Hash
	plugin.Logger(ctx).Warn("hydrateHash", "pw", pw, "result", result)
	return result, nil
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

func getPassword(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		panic(err)
	}

	quals := d.KeyColumnQuals

	hash := quals["hash"].GetStringValue()

	pwMatch, _, err := client.Passwords.GetExactPasswordBySHA1(hash)

	if err != nil {
		panic(err)
	}

	return pwMatch, nil
}

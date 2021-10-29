package hibp

import (
	"context"

	"github.com/turbot/hibp-go"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAccount() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_account",
		Description: "Breached accounts tracked by HIBP",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("account"),
			Hydrate:    listBreachedAccounts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getBreach,
		},
		Columns: append(breachColumns(), &plugin.Column{
			Name:        "account",
			Type:        proto.ColumnType_STRING,
			Description: "The email or phone account that was found in the paste (this field is required).",
			Transform:   transform.FromQual("account"),
		}),
	}
}

func listBreachedAccounts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)
	if err != nil {
		plugin.Logger(ctx).Error("hibp_account.listBreachedAccounts", "client.error", err)
		return nil, err
	}

	account := d.KeyColumnQuals["account"].GetStringValue()
	breaches, _, err := client.Breaches.ByAccount(account)
	if err != nil {
		plugin.Logger(ctx).Error("hibp_account.listBreachedAccounts", "api.error", err)
		return nil, err
	}

	for _, breach := range breaches {
		d.StreamListItem(ctx, breach)
	}
	return nil, nil
}

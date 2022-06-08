package hibp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"github.com/wneessen/go-hibp"
)

func tableAccount() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_account",
		Description: "Breached accounts tracked by HIBP",
		List: &plugin.ListConfig{
			Hydrate: listBreachedAccounts,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "account", Require: plugin.Required},
				{Name: "is_verified", Require: plugin.Optional},
				{Name: "domain", Require: plugin.Optional},
			},
			ShouldIgnoreError: ignore404Error,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			Hydrate:           getBreach,
			ShouldIgnoreError: ignore404Error,
		},
		Columns: append(breachColumns(), &plugin.Column{
			Name:        "account",
			Type:        proto.ColumnType_STRING,
			Description: "The email or phone account that was found in the breach (this field is required).",
			Transform:   transform.FromQual("account"),
		}),
	}
}

func listBreachedAccounts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getHibpClient(ctx, d)

	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals
	account := quals["account"].GetStringValue()

	requestOptions := []hibp.BreachOption{
		hibp.WithoutTruncate(),
	}

	if val, ok := d.KeyColumnQuals["is_verified"]; ok && val.GetBoolValue() {
		requestOptions = append(requestOptions, hibp.WithoutUnverified())
	}
	if val, ok := d.KeyColumnQuals["domain"]; ok && val.GetBoolValue() {
		requestOptions = append(requestOptions, hibp.WithDomain(val.GetStringValue()))
	}

	breaches, _, err := client.BreachApi.BreachedAccount(account, requestOptions...)

	if err != nil {
		return nil, err
	}

	for _, breach := range breaches {
		d.StreamListItem(ctx, breach)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

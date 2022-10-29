package hibp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/wneessen/go-hibp"
)

func tableHIBPBreach() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_breach",
		Description: "Breaches tracked by HIBP.",
		List: &plugin.ListConfig{
			Hydrate:    listBreaches,
			KeyColumns: plugin.OptionalColumns([]string{"is_verified", "domain"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getBreach,
		},
		Columns: breachColumns(),
	}
}

func listBreaches(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getHibpClient(ctx, d)

	if err != nil {
		return nil, err
	}

	requestOptions := []hibp.BreachOption{
		hibp.WithoutTruncate(),
	}

	if val, ok := d.KeyColumnQuals["is_verified"]; ok && val.GetBoolValue() {
		requestOptions = append(requestOptions, hibp.WithoutUnverified())
	}
	if val, ok := d.KeyColumnQuals["domain"]; ok && val.GetBoolValue() {
		requestOptions = append(requestOptions, hibp.WithDomain(val.GetStringValue()))
	}

	breaches, _, err := client.BreachAPI.Breaches(requestOptions...)

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

func getBreach(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getHibpClient(ctx, d)

	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals
	name := quals["name"].GetStringValue()

	breach, _, err := client.BreachAPI.BreachByName(name)
	if err != nil {
		return nil, err
	}

	return breach, nil
}

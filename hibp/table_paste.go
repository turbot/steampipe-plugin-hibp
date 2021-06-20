package hibp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
	"gitlab.com/wedtm/go-hibp"
)

func tablePaste() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_paste",
		Description: "Pastes tracked by HIBP",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("account"),
			Hydrate:    listPastes,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("account"),
			Hydrate:    getPaste,
		},
		Columns: []*plugin.Column{

			{Name: "account", Type: proto.ColumnType_STRING, Description: "The email account that was found in the paste (this field is required).", Transform: transform.FromField("account")},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "The paste service the record was retrieved from. Current values are: Pastebin, Pastie, Slexy, Ghostbin, QuickLeak, JustPaste, AdHocUrl, PermanentOptOut, OptOut"},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the paste as it was given at the source service. Combined with the 'source' attribute, this can be used to resolve the URL of the paste."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "The title of the paste as observed on the source site. This may be null and if so will be omitted from the response."},
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time that the paste was posted. This is taken directly from the paste site when this information is available but may be null if no date is published.."},
			{Name: "email_count", Type: proto.ColumnType_INT, Description: "The number of emails that were found when processing the paste."},
		},
	}
}

func listPastes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		panic(err)
	}

	quals := d.KeyColumnQuals
	pastes, _, err := client.Pastes.GetPastesByAccount(quals["account"].GetStringValue())

	if err != nil {
		panic(err)
	}

	for _, paste := range pastes {
		d.StreamListItem(ctx, paste)
	}
	return nil, nil
}

func getPaste(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		panic(err)
	}

	quals := d.KeyColumnQuals
	pastes, _, err := client.Pastes.GetPastesByAccount(quals["account"].GetStringValue())

	if err != nil {
		panic(err)
	}

	return pastes, nil
}

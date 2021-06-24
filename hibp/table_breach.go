package hibp

import (
	"context"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
	"gitlab.com/wedtm/go-hibp"
)

func tableBreach() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_breach",
		Description: "Breaches tracked by HIBP",
		List: &plugin.ListConfig{
			Hydrate: listBreaches,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name", "account"}),
			Hydrate:    getBreach,
		},
		Columns: []*plugin.Column{
			{Name: "account", Type: proto.ColumnType_STRING, Description: "The email account that was found in the paste (this field is required).", Transform: transform.FromValue().NullIfZero(), Hydrate: getAccountBreaches},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "A Pascal-cased name representing the breach which is unique across all other breaches. This value never changes and may be used to name dependent assets (such as images) but should not be shown directly to end users (see the 'title' field instead)."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "A descriptive title for the breach suitable for displaying to end users. It's unique across all breaches but individual values may change in the future (i.e. if another breach occurs against an organisation already in the system). If a stable value is required to reference the breach, refer to the 'name' field instead."},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "The domain of the primary website the breach occurred on. This may be used for identifying other assets external systems may have for the site."},
			{Name: "breach_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date (with no time) the breach originally occurred on in ISO 8601 format. This is not always accurate — frequently breaches are discovered and reported long after the original incident. Use this field as a guide only."},
			{Name: "added_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time (precision to the minute) the breach was added to the system."},
			{Name: "modified_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time (precision to the minute) the breach was modified. This will only differ from the added_date attribute if other attributes represented here are changed or data in the breach itself is changed (i.e. additional data is identified and loaded). It is always either equal to or greater then the added_date field, never less than."},
			{Name: "pwn_count", Type: proto.ColumnType_INT, Description: "The total number of accounts loaded into the system. This is usually less than the total number reported by the media due to duplication or other data integrity issues in the source data."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Contains an overview of the breach represented in HTML markup. The description may include markup such as emphasis and strong tags as well as hyperlinks."},
			{Name: "data_classes", Type: proto.ColumnType_JSON, Description: "This field describes the nature of the data compromised in the breach and contains an array of impacted data classes."},
			{Name: "is_verified", Type: proto.ColumnType_BOOL, Description: "Indicates that the breach is considered unverified. An unverified breach may not have been hacked from the indicated website. An unverified breach is still loaded into HIBP when there's sufficient confidence that a significant portion of the data is legitimate."},
			{Name: "is_fabricated", Type: proto.ColumnType_BOOL, Description: "Indicates that the breach is considered fabricated. A fabricated breach is unlikely to have been hacked from the indicated website and usually contains a large amount of manufactured data. However, it still contains legitimate email addresses and asserts that the account owners were compromised in the alleged breach."},
			{Name: "is_sensitive", Type: proto.ColumnType_BOOL, Description: "Indicates if the breach is considered sensitive. The public API will not return any accounts for a breach flagged as sensitive."},
			{Name: "is_retired", Type: proto.ColumnType_BOOL, Description: "Indicates if the breach has been retired. This data has been permanently removed and will not be returned by the API."},
			{Name: "is_spam_list", Type: proto.ColumnType_BOOL, Description: "Indicates if the breach is considered a spam list. This flag has no impact on any other attributes but it means that the data has not come as a result of a security compromise."},
			{Name: "logo_path", Type: proto.ColumnType_STRING, Description: "A URI that specifies where a logo for the breached service can be found. Logos are always in PNG format."},
		},
	}
}

func listBreaches(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		return nil, err
	}

	breaches, _, err := client.Breaches.ListBreaches()

	if err != nil {
		return nil, err
	}

	for _, breach := range breaches {
		d.StreamListItem(ctx, breach)
	}
	return nil, nil
}

func getBreach(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals
	name := quals["name"].GetStringValue()

	breaches, _, err := client.Breaches.GetBreach(name)

	if err != nil {
		return nil, err
	}

	return breaches, nil
}

func getAccountBreaches(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := hibp.NewClient(*GetConfig(d.Connection).ApiKey, nil)

	if err != nil {
		return nil, err
	}

	quals := d.KeyColumnQuals
	account := quals["account"].GetStringValue()

	if account == "" {
		return nil, nil
	}

	breaches, _, err := client.Breaches.ByAccount(account)

	log.Printf("%v", breaches)
	if err != nil {
		return nil, err
	}

	return breaches, nil
}

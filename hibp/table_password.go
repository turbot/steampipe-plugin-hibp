package hibp

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type psswrdRow struct {
	Plain *string
	Hash  *string
	Count *int64
}

func tablePassword() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_password",
		Description: "Password (hashes) tracked by HIBP",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AnyColumn([]string{"plain", "hash"}),
			Hydrate:    listPasswords,
		},
		Columns: []*plugin.Column{
			{Name: "plain", Type: proto.ColumnType_STRING, Description: "The plain-text of the compromised password (sent as a hash to the API)."},
			{Name: "hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

func listPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getHibpClient(ctx, d)

	if err != nil {
		return nil, err
	}

	var plain *string

	hash := ""
	if h, ok := d.KeyColumnQuals["hash"]; ok {
		hash = h.GetStringValue()
	} else {
		p := d.KeyColumnQuals["plain"].GetStringValue()
		plain = &p
		hash = fmt.Sprintf("%x", sha1.Sum([]byte(*plain)))
	}

	if len(hash) < 40 {
		return nil, fmt.Errorf("password hash needs to be a SHA1 digest (40 character hexadecimal)")
	}

	match, _, err := client.PwnedPassApi.CheckSHA1(hash)

	if err != nil {
		return nil, err
	}

	row := &psswrdRow{
		Hash:  &match.Hash,
		Plain: plain,
		Count: &match.Count,
	}

	d.StreamListItem(ctx, row)

	return nil, nil
}

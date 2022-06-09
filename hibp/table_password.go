package hibp

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

type psswrdRow struct {
	PasswordHash string
	Count        int64
}

func tablePassword() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_password",
		Description: "Password (hashes) tracked by HIBP",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AnyColumn([]string{"password", "password_hash"}),
			Hydrate:    listPasswords,
		},
		Columns: []*plugin.Column{
			{Name: "password", Type: proto.ColumnType_STRING, Transform: transform.FromQual("password").NullIfZero(), Description: "The plain-text of the compromised password (sent as a hash to the API)."},
			{Name: "password_hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

func listPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getHibpClient(ctx, d)

	if err != nil {
		return nil, err
	}

	var password string

	passwordHash := ""
	if h := d.KeyColumnQualString("password_hash"); len(h) > 0 {
		passwordHash = h
	} else {
		password = d.KeyColumnQualString("password")
		passwordHash = fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	}

	if len(passwordHash) < 40 {
		return nil, fmt.Errorf("password hash needs to be a SHA1 digest (40 character hexadecimal)")
	}

	match, _, err := client.PwnedPassApi.CheckSHA1(passwordHash)
	if err != nil {
		return nil, err
	}

	if match != nil {
		row := &psswrdRow{
			PasswordHash: match.Hash,
			Count:        match.Count,
		}

		d.StreamListItem(ctx, row)
	}

	return nil, nil
}

package hibp

import (
	"context"
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
	"github.com/wneessen/go-hibp"
)

type psswrdRow struct {
	hibp.Match
	HashPrefix string
}

func tableHIBPPassword() *plugin.Table {
	return &plugin.Table{
		Name:        "hibp_password",
		Description: "Password (hashes) tracked by HIBP.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AnyColumn([]string{"plaintext", "hash", "hash_prefix"}),
			Hydrate:    listPasswords,
		},
		Columns: []*plugin.Column{
			{Name: "plaintext", Type: proto.ColumnType_STRING, Transform: transform.FromQual("plaintext").NullIfZero(), Description: "The plain-text of the compromised password (sent as a hash to the API)."},
			{Name: "hash", Type: proto.ColumnType_STRING, Description: "The hash of the compromised password."},
			{Name: "hash_prefix", Type: proto.ColumnType_STRING, Description: "The first 5-char prefix of the hash of the compromised password."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "The total number of times this password has been found compromised."},
		},
	}
}

func listPasswords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getHibpClient(ctx, d)

	if err != nil {
		return nil, err
	}

	hashPrefixToSearch := ""
	if hash := d.KeyColumnQualString("hash"); len(hash) > 0 {
		if len(hash) != 40 {
			return nil, fmt.Errorf("'hash' must be exactly 40 characters hexadecimal")
		}
		// make sure that this is a valid hex string
		_, err := strconv.ParseUint(hash, 16, 64)
		if err != nil {
			return nil, fmt.Errorf("'hash' is not a valid SHA-1 hash")
		}
		hashPrefixToSearch = hash
	} else if prefix := d.KeyColumnQualString("hash_prefix"); len(prefix) > 0 {
		hashPrefixToSearch = d.KeyColumnQualString("hash_prefix")
		if len(hashPrefixToSearch) < 5 {
			return nil, fmt.Errorf("'hash_prefix' must be at least 5 characters")
		}

		// make sure that this is a valid hex string
		_, err := strconv.ParseUint(hashPrefixToSearch, 16, 64)
		if err != nil {
			return nil, fmt.Errorf("'hash_prefix' is not a valid SHA-1 hash prefix")
		}
	} else {
		hashPrefixToSearch = fmt.Sprintf("%x", sha1.Sum([]byte(d.KeyColumnQualString("plaintext"))))
	}

	matches, _, err := client.PwnedPassApi.ListHashesPrefix(hashPrefixToSearch[:5])
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		if match.Count > 0 && strings.HasPrefix(match.Hash, hashPrefixToSearch) {
			row := &psswrdRow{
				Match:      match,
				HashPrefix: match.Hash[:len(hashPrefixToSearch)],
			}

			d.StreamListItem(ctx, row)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) < 1 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

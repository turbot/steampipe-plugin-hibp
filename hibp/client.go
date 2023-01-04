package hibp

import (
	"context"
	"os"
	"steampipe-plugin-hibp/constants"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/wneessen/go-hibp"
)

func getHibpClient(ctx context.Context, d *plugin.QueryData) (*hibp.Client, error) {
	// Try to load client from cache
	if cachedData, ok := d.ConnectionManager.Cache.Get(constants.CacheKeyHibpClient); ok {
		return cachedData.(*hibp.Client), nil
	}

	// Get apiKey
	apiKey, err := getKeysFromConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// create the hibp client
	client := createClient(ctx, apiKey)

	// save client in cache
	d.ConnectionManager.Cache.Set(constants.CacheKeyHibpClient, client)

	return client, nil
}

// getKeysFromConfig fetches the apiKey from the connection config
// falls back to the environment variables if it cannot find one in the config
// returns an error if api key could not be resolved
func getKeysFromConfig(ctx context.Context, d *plugin.QueryData) (apiKey string, _ error) {
	config := GetConfig(d.Connection)

	// Get the authorization publicKey
	apiKey = os.Getenv(constants.EnvKeyApiKey)
	if config.ApiKey != nil {
		apiKey = *config.ApiKey
	}

	// Return nil since some tables like `hibp_password` and `hibp_breach` don't need an API key
	if len(apiKey) == 0 {
		return "", nil
	}

	return apiKey, nil
}

func createClient(ctx context.Context, apiKey string) *hibp.Client {
	clientOptions := []hibp.Option{
		hibp.WithAPIKey(apiKey),
		hibp.WithUserAgent("Turbot Steampipe (+https://steampipe.io)"),
		hibp.WithHTTPTimeout(3 * time.Second),
		hibp.WithRateLimitSleep(),
	}

	cl := hibp.New(clientOptions...)
	return &cl
}


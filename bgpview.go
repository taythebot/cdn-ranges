package cdn_ranges

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"go.uber.org/ratelimit"
)

var (
	rateLimiterOnce sync.Once
	rateLimiter     ratelimit.Limiter
)

type IPGuideResponse struct {
	Asn          int    `json:"asn"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	Country      string `json:"country"`
	Rir          string `json:"rir"`
	Routes       struct {
		V4 []string `json:"v4"`
		V6 []string `json:"v6"`
	} `json:"routes"`
}

// GetRateLimit returns a global reusable rate limiter
func GetRateLimit() ratelimit.Limiter {
	rateLimiterOnce.Do(func() {
		rateLimiter = ratelimit.New(3)
	})

	return rateLimiter
}

// ASNPrefixes fetches IPv4 and IPv6 prefixes for a given ASN
func ASNPrefixes(ctx context.Context, asn int) ([]string, []string, error) {
	GetRateLimit().Take()

	resp, err := HttpGet(ctx, fmt.Sprintf("https://ip.guide/AS%d", asn))
	if err != nil {
		return nil, nil, err
	}

	var response IPGuideResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println(string(resp))
		return nil, nil, fmt.Errorf("failed to perform unmarshal json: %w", err)
	}

	return response.Routes.V4, response.Routes.V6, nil
}

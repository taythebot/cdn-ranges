package provider

import (
	"context"
	"encoding/json"
	"fmt"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*GCore)(nil)

type GCore struct{}

type GCoreResponse struct {
	Addresses     []string `json:"addresses"`
	Ipv6Addresses []string `json:"ipv6_addresses"`
}

func (p *GCore) Name() string {
	return "Gcore"
}

func (p *GCore) Fetch(ctx context.Context) ([]string, []string, error) {
	resp, err := cdn_ranges.HttpGet(ctx, "https://api.gcore.com/cdn/public-ip-list")
	if err != nil {
		return nil, nil, err
	}

	var response FastlyResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return response.Addresses, response.Ipv6Addresses, nil
}

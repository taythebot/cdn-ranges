package provider

import (
	"context"
	"encoding/json"
	"fmt"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Fastly)(nil)

type Fastly struct{}

type FastlyResponse struct {
	Addresses     []string `json:"addresses"`
	Ipv6Addresses []string `json:"ipv6_addresses"`
}

func (p *Fastly) Name() string {
	return "Fastly"
}

func (p *Fastly) Fetch(ctx context.Context) ([]string, []string, error) {
	resp, err := cdn_ranges.HttpGet(ctx, "https://api.fastly.com/public-ip-list")
	if err != nil {
		return nil, nil, err
	}

	var response FastlyResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return response.Addresses, response.Ipv6Addresses, nil
}

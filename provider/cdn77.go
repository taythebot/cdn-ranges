package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*CDN77)(nil)

type CDN77 struct{}

type CDN77Response struct {
	UpdatedAt time.Time `json:"updated_at"`
	Prefixes  []struct {
		Prefix string `json:"prefix"`
	} `json:"prefixes"`
}

func (p *CDN77) Name() string {
	return "CDN77"
}

func (p *CDN77) Fetch(ctx context.Context) ([]string, []string, error) {
	resp, err := cdn_ranges.HttpGet(ctx, "https://prefixlists.tools.cdn77.com/public_lmax_prefixes.json")
	if err != nil {
		return nil, nil, err
	}

	var response CDN77Response
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	v4 := make([]string, 0, len(response.Prefixes))
	for _, prefix := range response.Prefixes {
		v4 = append(v4, prefix.Prefix)
	}

	return v4, nil, nil
}

package provider

import (
	"context"
	"encoding/json"
	"fmt"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Cloudfront)(nil)

type Cloudfront struct{}

type CloudfrontResponse struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix           string `json:"ip_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"prefixes"`
	Ipv6Prefixes []struct {
		Ipv6Prefix         string `json:"ipv6_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"ipv6_prefixes"`
}

func (p *Cloudfront) Name() string {
	return "Cloudfront"
}

func (p *Cloudfront) Fetch(ctx context.Context) ([]string, []string, error) {
	resp, err := cdn_ranges.HttpGet(ctx, "https://ip-ranges.amazonaws.com/ip-ranges.json")
	if err != nil {
		return nil, nil, err
	}

	var response CloudfrontResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	var v4 []string
	for _, prefix := range response.Prefixes {
		if prefix.Service == "CLOUDFRONT" {
			v4 = append(v4, prefix.IPPrefix)
		}
	}

	var v6 []string
	for _, prefix := range response.Ipv6Prefixes {
		if prefix.Service == "CLOUDFRONT" {
			v4 = append(v4, prefix.Ipv6Prefix)
		}
	}

	return v4, v6, nil
}

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Medianova)(nil)

type Medianova struct{}

type MedianovaResponse struct {
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   struct {
		Ipv4Cidrs []string `json:"ipv4_cidrs"`
		Ipv6Cidrs []string `json:"ipv6_cidrs"`
	} `json:"result"`
	Success bool `json:"success"`
}

func (p *Medianova) Name() string {
	return "Medianova"
}

func (p *Medianova) Fetch(ctx context.Context) ([]string, []string, error) {
	resp, err := cdn_ranges.HttpGet(ctx, "https://cloud.medianova.com/api/v1/ip/blocks-list")
	if err != nil {
		return nil, nil, err
	}

	var response MedianovaResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	if !response.Success {
		return nil, nil, errors.New("status is not ok")
	}

	return response.Result.Ipv4Cidrs, response.Result.Ipv6Cidrs, nil
}

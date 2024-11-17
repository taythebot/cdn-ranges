package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Imperva)(nil)

type Imperva struct{}

type ImpervaResponse struct {
	IPRanges   []string `json:"ipRanges"`
	Ipv6Ranges []string `json:"ipv6Ranges"`
	Res        int      `json:"res"`
	ResMessage string   `json:"res_message"`
	DebugInfo  struct {
		IDInfo string `json:"id-info"`
	} `json:"debug_info"`
}

func (p *Imperva) Name() string {
	return "Imperva"
}

func (p *Imperva) Fetch(ctx context.Context) ([]string, []string, error) {
	resp, err := cdn_ranges.HttpGet(ctx, "https://my.imperva.com/api/integration/v1/ips")
	if err != nil {
		return nil, nil, err
	}

	var response ImpervaResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	if response.ResMessage != "OK" {
		return nil, nil, errors.New("status is not ok")
	}

	return response.IPRanges, response.Ipv6Ranges, nil
}

package provider

import (
	"context"
	"strings"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*CacheFly)(nil)

type CacheFly struct{}

func (p *CacheFly) Name() string {
	return "CacheFly"
}

func (p *CacheFly) Fetch(ctx context.Context) ([]string, []string, error) {
	resp, err := cdn_ranges.HttpGet(ctx, "https://cachefly.cachefly.net/ips/rproxy.txt")
	if err != nil {
		return nil, nil, err
	}

	v4 := strings.Split(string(resp), "\n")
	return v4[:len(v4)-1], nil, nil
}

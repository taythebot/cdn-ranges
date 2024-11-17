package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*CDNetworks)(nil)

type CDNetworks struct{}

func (p *CDNetworks) Name() string {
	return "CDNetworks"
}

func (p *CDNetworks) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 36408)
}

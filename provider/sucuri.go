package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Sucuri)(nil)

type Sucuri struct{}

func (p *Sucuri) Name() string {
	return "Sucuri"
}

func (p *Sucuri) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 30148)
}

package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Bunny)(nil)

type Bunny struct{}

func (p *Bunny) Name() string {
	return "Bunny"
}

func (p *Bunny) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 200325)
}

package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Edgecast)(nil)

type Edgecast struct{}

func (p *Edgecast) Name() string {
	return "Edgecast"
}

func (p *Edgecast) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 15133)
}

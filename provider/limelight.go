package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Limelight)(nil)

type Limelight struct{}

func (p *Limelight) Name() string {
	return "Limelight"
}

func (p *Limelight) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 22822)
}

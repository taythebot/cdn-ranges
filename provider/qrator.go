package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Qrator)(nil)

type Qrator struct{}

func (p *Qrator) Name() string {
	return "Qrator"
}

func (p *Qrator) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 200449)
}

package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*X4B)(nil)

type X4B struct{}

func (p *X4B) Name() string {
	return "X4B"
}

func (p *X4B) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 136165)
}

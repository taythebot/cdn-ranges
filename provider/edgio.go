package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Edgio)(nil)

type Edgio struct{}

func (p *Edgio) Name() string {
	return "Edgio"
}

func (p *Edgio) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 60261)
}

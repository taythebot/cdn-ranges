package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*DDoSGuard)(nil)

type DDoSGuard struct{}

func (p *DDoSGuard) Name() string {
	return "DDoS-Guard"
}

func (p *DDoSGuard) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 57724)
}

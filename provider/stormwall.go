package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*StormWall)(nil)

type StormWall struct{}

func (p *StormWall) Name() string {
	return "StormWall"
}

func (p *StormWall) Fetch(ctx context.Context) ([]string, []string, error) {
	return cdn_ranges.ASNPrefixes(ctx, 59796)
}

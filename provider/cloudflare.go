package provider

import (
	"context"
	"strings"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Cloudflare)(nil)

type Cloudflare struct{}

func (p *Cloudflare) Name() string {
	return "Cloudflare"
}

func (p *Cloudflare) Fetch(ctx context.Context) ([]string, []string, error) {
	respV4, err := cdn_ranges.HttpGet(ctx, "https://www.cloudflare.com/ips-v4")
	if err != nil {
		return nil, nil, err
	}
	v4 := strings.Split(string(respV4), "\n")

	respV6, err := cdn_ranges.HttpGet(ctx, "https://www.cloudflare.com/ips-v6")
	if err != nil {
		return nil, nil, err
	}
	v6 := strings.Split(string(respV6), "\n")

	return v4[:len(v4)-1], v4[:len(v6)-1], nil
}

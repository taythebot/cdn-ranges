package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Akamai)(nil)

type Akamai struct{}

func (p *Akamai) Name() string {
	return "Akamai"
}

func (p *Akamai) Fetch(ctx context.Context) ([]string, []string, error) {
	asn1_v4, asn1_v6, err := cdn_ranges.ASNPrefixes(ctx, 12222)
	if err != nil {
		return nil, nil, err
	}

	asn2_v4, asn2_v6, err := cdn_ranges.ASNPrefixes(ctx, 16625)
	if err != nil {
		return nil, nil, err
	}

	return append(asn1_v4, asn2_v4...), append(asn1_v6, asn2_v6...), nil
}

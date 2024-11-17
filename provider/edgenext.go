package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*EdgeNext)(nil)

type EdgeNext struct{}

var EdgeNextASNs = []int{
	139057,
	149981,
}

func (p *EdgeNext) Name() string {
	return "EdgeNext"
}

func (p *EdgeNext) Fetch(ctx context.Context) ([]string, []string, error) {
	var (
		v4 []string
		v6 []string
	)
	for _, asn := range EdgeNextASNs {
		asn_v4, asn_v6, err := cdn_ranges.ASNPrefixes(ctx, asn)
		if err != nil {
			return nil, nil, err
		}

		v4 = append(v4, asn_v4...)
		v6 = append(v6, asn_v6...)
	}

	return v4, v6, nil
}

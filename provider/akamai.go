package provider

import (
	"context"

	cdn_ranges "github.com/taythebot/cdn-ranges"
)

var _ Provider = (*Akamai)(nil)

type Akamai struct{}

var AkamaiAsns = []int{
	12222,
	16625,
	16702,
	17204,
	18680,
	18717,
	20189,
	20940,
	21342,
	21357,
	21399,
	22207,
	22452,
	23454,
	23455,
	23903,
	24319,
	26008,
	30675,
	31107,
	31108,
	31109,
	31110,
	31377,
	33047,
	33905,
	34164,
	34850,
	35204,
	35993,
	35994,
	36183,
	39836,
	43639,
	55409,
	55770,
	63949,
	133103,
	393560,
}

func (p *Akamai) Name() string {
	return "Akamai"
}

func (p *Akamai) Fetch(ctx context.Context) ([]string, []string, error) {
	var (
		v4 []string
		v6 []string
	)
	for _, asn := range AkamaiAsns {
		asn_v4, asn_v6, err := cdn_ranges.ASNPrefixes(ctx, asn)
		if err != nil {
			return nil, nil, err
		}

		v4 = append(v4, asn_v4...)
		v6 = append(v6, asn_v6...)
	}

	return v4, v6, nil
}

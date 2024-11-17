package provider

import (
	"context"
)

var _ Provider = (*Leaseweb)(nil)

type Leaseweb struct{}

func (p *Leaseweb) Name() string {
	return "Leaseweb"
}

func (p *Leaseweb) Fetch(_ context.Context) ([]string, []string, error) {
	v4 := []string{
		"185.28.68.0/24",
		"185.28.69.0/24",
		"185.28.70.0/24",
		"185.28.71.0/24",
		"89.255.254.0/24",
		"89.255.255.0/24",
		"89.255.248.0/24",
		"89.255.249.0/24",
		"89.255.250.0/24",
		"89.255.251.0/24",
		"89.255.252.0/24",
		"67.208.209.176/28",
	}

	v6 := []string{
		"2a00:9d20:1::/48",
		"2a00:9d20:2::/48",
		"2a00:9d20:3::/48",
		"2a00:9d20:4::/48",
		"2a00:9d20:5::/48",
		"2a00:9d20:6::/48",
		"2a00:9d20:53::/48",
		"2a00:9d20:5353::/48",
		"2a00:9d20:7::/48",
	}

	return v4, v6, nil
}

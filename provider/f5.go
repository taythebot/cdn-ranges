package provider

import (
	"context"
)

var _ Provider = (*F5)(nil)

type F5 struct{}

func (p *F5) Name() string {
	return "F5"
}

func (p *F5) Fetch(_ context.Context) ([]string, []string, error) {
	v4 := []string{
		"159.60.188.0/24",
		"159.60.189.0/24",
		"159.60.190.0/24",
		"159.60.191.0/24",
		"159.60.187.0/24",
	}

	return v4, nil, nil
}

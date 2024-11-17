package provider

import (
	"context"
)

// Provider interface for all providers
type Provider interface {
	Name() string
	Fetch(ctx context.Context) (ipv4 []string, ipv6 []string, err error)
}

var Providers = []Provider{
	&Akamai{},
	&Bunny{},
	&CacheFly{},
	&CDNetworks{},
	&Cloudflare{},
	&Cloudfront{},
	&DDoSGuard{},
	&Edgecast{},
	&EdgeNext{},
	&Edgio{},
	&F5{},
	&Fastly{},
	&GCore{},
	&Imperva{},
	&Leaseweb{},
	&Limelight{},
	&Medianova{},
	&Qrator{},
	&StormWall{},
	&Sucuri{},
	&X4B{},
}

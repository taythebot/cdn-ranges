package main

// Provider ...
type Provider struct {
	ASNs     []string
	IPv4List string
	IPv6List string
}

// ProvidersMap lists all providers available
var ProvidersMap = map[string]Provider{
	"akamai": {
		ASNs: []string{"12222", "20940"},
	},
	"cloudflare": {
		IPv4List: "https://www.cloudflare.com/ips-v4",
		IPv6List: "https://www.cloudflare.com/ips-v6",
	},
}

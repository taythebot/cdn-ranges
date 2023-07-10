package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

// Options for CLI
type Options struct {
	Format    string
	Output    string
	Providers []string
	IPv4      bool
	IPv6      bool
}

func fatal(msg string, v ...interface{}) {
	fmt.Println("[*] Exiting Program: " + fmt.Sprintf(msg, v...))
	os.Exit(1)
}

func main() {
	options := &Options{}
	flag.StringVarP(&options.Format, "format", "f", "txt", "output format (txt, csv, json)")
	flag.StringVarP(&options.Output, "output", "o", "", "Output file")
	flag.StringArrayVarP(&options.Providers, "providers", "p", nil, "Providers to query")
	flag.BoolVarP(&options.IPv6, "ipv4", "4", false, "Download IPv4 addresses")
	flag.BoolVarP(&options.IPv6, "ipv6", "6", false, "Download IPv6 addresses")
	flag.Parse()

	if options.Format != "txt" && options.Format != "csv" && options.Format != "json" {
		fatal("Invalid format provided. Must be one of txt, csv, json")
	}

	// Parse providers
	var providers []Provider
	if len(options.Providers) > 0 {
		for _, op := range options.Providers {
			p, ok := ProvidersMap[op]
			if !ok {
				fatal("Invalid provider '%s'", op)
				os.Exit(1)
			}

			providers = append(providers, p)
		}
	} else {
		for _, p := range ProvidersMap {
			providers = append(providers, p)
		}
	}

	if len(providers) == 0 {
		fatal("No valid providers specified")
	}

	// Check IP types
	allIPs := (options.IPv4 && options.IPv6) || (!options.IPv4 && !options.IPv6)
	if allIPs {
		fmt.Printf("[*] Downloading IPv4 and IPv6 ranges from %d providers\n", len(providers))
	} else if options.IPv4 {
		fmt.Printf("[*] Downloading IPv4 ranges from %d providers\n", len(providers))
	} else if options.IPv6 {
		fmt.Printf("[*] Downloading IPv6 ranges from %d providers\n", len(providers))
	}

	// Download ip ranges
	bgpview := &http.Client{
		Transport: &http.Transport{MaxIdleConnsPerHost: 20},
		Timeout:   60 * time.Second,
	}

	for _, provider := range providers {
		var ipv4, ipv6 []string

		if len(provider.ASNs) > 0 {
			v4, v6, err := FetchBGPView(bgpview, provider.ASNs)
			if err != nil {
				fatal(err.Error())
			}

			ipv4 = v4
			ipv6 = v6
		}

		if provider.IPv4List != "" {
		}

		fmt.Println(ipv4)
		fmt.Println(ipv6)
	}
}

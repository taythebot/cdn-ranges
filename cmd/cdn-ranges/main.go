package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/taythebot/cdn-ranges/provider"
)

var formats = []string{"txt", "csv", "json"}

type Output struct {
	Provider string `json:"provider"`
	Type     string `json:"type"`
	Range    string `json:"range"`
}

func main() {
	outputFlag := flag.String("output", "ranges.txt", "output file name")
	formatFlag := flag.String("format", "txt", "output format (txt, csv, json)")
	providerFlag := flag.String("provider", "", "provider name in lowercase")
	v4Flag := flag.Bool("v4", false, "Fetch IPv4 ranges only")
	v6Flag := flag.Bool("v6", false, "Fetch IPv6 ranges only")
	flag.Parse()

	if !slices.Contains(formats, *formatFlag) {
		fmt.Println("[Fatal] Format must be one of (txt, csv, json)")
		os.Exit(1)
	}

	providers := provider.Providers
	if *providerFlag != "" {
		var valid bool
		for _, p := range provider.Providers {
			if strings.ToLower(p.Name()) == strings.ToLower(*providerFlag) {
				valid = true
				providers = []provider.Provider{p}
				break
			}
		}

		if !valid {
			fmt.Println("[Fatal] Invalid provider. Check documentation for list of available providers")
			os.Exit(1)
		}
	}

	if *outputFlag == "" {
		fmt.Println("[Fatal] Output file must be specified (-output <file>)")
		os.Exit(1)
	}

	if *v4Flag {
		fmt.Println("[Info] Fetching only IPv4 ranges")
	} else if *v6Flag {
		fmt.Println("[Info] Fetching only IPv6 ranges")
	} else {
		fmt.Println("[Info] Fetching IPv4 and IPv6 ranges")
	}

	outputFile, err := os.Create(*outputFlag)
	if err != nil {
		fmt.Printf("[Fatal] Failed to create output file: %w\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	var (
		count  int // No atomic usage since writer thread is 1
		output = make(chan Output, 1)
		queue  = make(chan provider.Provider, 4)
		g, ctx = errgroup.WithContext(context.Background())
		wg     sync.WaitGroup
	)

	// Output file write
	wg.Add(1)
	go func() {
		defer wg.Done()

		var csvWriter *csv.Writer
		if *formatFlag == "csv" {
			csvWriter = csv.NewWriter(outputFile)
			csvWriter.Write([]string{"provider", "type", "range"})
		}

		for o := range output {
			switch *formatFlag {
			case "txt":
				outputFile.WriteString(o.Range + "\n")
			case "json":
				j, _ := json.Marshal(o)
				outputFile.WriteString(string(j) + "\n")
			case "csv":
				csvWriter.Write([]string{o.Provider, o.Type, o.Range})
			}

			count++
		}

		if *formatFlag == "csv" {
			csvWriter.Flush()
		}
	}()

	// Task queue for providers
	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, p := range providers {
			queue <- p
		}

		close(queue)
	}()

	// Workers
	for i := 0; i < 4; i++ {
		g.Go(func() error {
			for p := range queue {
				fmt.Printf("[Info] Fetching %s ranges\n", p.Name())

				v4, v6, err := p.Fetch(ctx)
				if err != nil {
					return fmt.Errorf("failed to fetch %s ranges: %w", p.Name(), err)
				}

				if *v4Flag || (!*v4Flag && !*v6Flag) {
					for _, r := range v4 {
						output <- Output{
							Provider: p.Name(),
							Type:     "ipv4",
							Range:    r,
						}
					}
				}

				if *v6Flag || (!*v4Flag && !*v6Flag) {
					for _, r := range v6 {
						output <- Output{
							Provider: p.Name(),
							Type:     "ipv6",
							Range:    r,
						}
					}
				}
			}

			return nil
		})
	}

	// Wait for data fetch to finish
	if err := g.Wait(); err != nil {
		fmt.Printf("[Fatal] Failed to fetch ranges: %w\n", err)
		os.Exit(1)
	}

	// Wait for all writes to finish
	close(output)
	wg.Wait()

	fmt.Printf("[Info] Successfully downloaded %d ranges in %s\n", count, *outputFlag)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BGPViewResponse from API
type BGPViewResponse struct {
	Status string `json:"status"`
	Data   struct {
		Ipv4Prefixes []struct {
			Prefix string `json:"prefix"`
		} `json:"ipv4_prefixes"`
		Ipv6Prefixes []struct {
			Prefix string `json:"prefix"`
		} `json:"ipv6_prefixes"`
	} `json:"data"`
}

// FetchBGPView fetches IP ranges from BGPView API
func FetchBGPView(client *http.Client, asns []string) ([]string, []string, error) {
	var (
		ipv4 []string
		ipv6 []string
	)

	for _, asn := range asns {
		req, err := http.NewRequest("GET", "https://api.bgpview.io/asn/"+asn+"/prefixes", nil)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create BGPView request for ASN '%s': %s", asn, err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to fetch BGPView data for ASN '%s': %s", asn, err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read BGPView response body for ASN '%s': %s", asn, err)
		}

		resp.Body.Close()

		var data BGPViewResponse
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal BGPView response body for ASN '%s': %s", asn, err)
		}

		for _, prefix := range data.Data.Ipv4Prefixes {
			ipv4 = append(ipv4, prefix.Prefix)
		}

		for _, prefix := range data.Data.Ipv6Prefixes {
			ipv6 = append(ipv6, prefix.Prefix)
		}

		time.Sleep(1000)
	}

	return ipv4, ipv6, nil
}

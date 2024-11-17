package cdn_ranges

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	httpClientOnce sync.Once
	httpClient     *http.Client
)

// GetHttpClient returns a global reusable HTTP client
func GetHttpClient() *http.Client {
	httpClientOnce.Do(func() {
		httpClient = &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 20,
			},
			Timeout: 10 * time.Second,
		}
	})

	return httpClient
}

// HttpGet performs an HTTP get request with context
func HttpGet(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req = req.WithContext(ctx)

	resp, err := GetHttpClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP body: %w", err)
	}

	return body, nil
}

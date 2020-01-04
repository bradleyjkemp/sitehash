package sitehash

import (
	"net/http"
	"net/url"
	"sort"
)

func getHeaders(url *url.URL) (string, []string, error) {
	// Use GET instead of HEAD to maximise compatibility
	resp, err := http.Get(url.String())
	if err != nil {
		return "", nil, err
	}
	// Discard body
	resp.Body.Close()

	headers := make([]string, 0, len(resp.Header))
	for header := range resp.Header {
		headers = append(headers, header)
	}
	// Try to ensure headers are deterministic
	sort.Slice(headers, func(i, j int) bool {
		return headers[i] < headers[j]
	})

	return resp.Status, headers, nil
}

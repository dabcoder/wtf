package utils

import (
	"net/http"
)

// HTTPClient takes a Url and additional parameters, makes a GET request to said Url, and returns an http response.
func HTTPClient(url string, apiKey string, value string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resp.Header.Add("Accept", "application/json")
	resp.Header.Add("Content-Type", "application/json")

	if apiKey != "" && value != "" {
		resp.Header.Add(apiKey, value)
	}

	return resp, nil
}

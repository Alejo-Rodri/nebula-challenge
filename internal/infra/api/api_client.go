package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func parseJSON(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func (c *ApiClient) Info() (ApiInfoResponse, error) {
	var result ApiInfoResponse
	
	resp, err := c.http.Get(c.baseURL + "/info")
	if err != nil {
		return result, fmt.Errorf("GET /info: %w, %s", ErrConnection, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr ApiErrorsResponse
		if err := parseJSON(resp.Body, &apiErr); err != nil {
			return result, fmt.Errorf("GET /info decode error body: %w", err)
		}

		return result, fmt.Errorf("%w: %s", ErrInvalidRequest, apiErr)
	}

	
	if err := parseJSON(resp.Body, &result); err != nil {
		return result, fmt.Errorf("GET /info decode body: %w", ErrInvalidResponse)
	}

	return result, nil
}

func (c *ApiClient) Analyze(host string) (ApiAnalyzeResponse, error) {
	var result ApiAnalyzeResponse
	return result, nil
}
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func get[T any](c *ApiClient, endpoint string, query url.Values) (T, error) {
	var result T

	baseURL, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return result, printError("GET", endpoint, ErrParsingUrlRequest, err)
	}

	if query == nil {
		baseURL.RawQuery = ""
	} else {
		baseURL.RawQuery = query.Encode()
	}

	fmt.Printf("url: %s\n", baseURL.String())
	resp, err := c.http.Get(baseURL.String())
	if err != nil {
		return result, printError("GET", endpoint, ErrConnection, err)
	}

	defer resp.Body.Close()

	fmt.Printf("status code: %d\n", resp.StatusCode)
	if err := validateResponse(resp); err != nil {
		return result, err
	}

	if err := parseJSON(resp.Body, &result); err != nil {
		return result, printError("GET", endpoint, ErrInvalidResponse, err)
	}

	return result, nil
}

func parseJSON(r io.Reader, v any) error {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	return json.NewDecoder(r).Decode(v)
}

func printError(verb, endpoint string, errorType, errorThrown error) error {
	return fmt.Errorf(verb + " " + endpoint + ": %w, %s", errorType, errorThrown)
}

// handles only wrong use cases
func validateResponse(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return printError("", "UNKNOWN", ErrReadingBodyResponse, err)
	}

	var apiErr ApiErrorsResponse
	if json.Unmarshal(body, &apiErr) == nil {
		return printError("", "UNKNOWN", ErrInvalidRequest, err)
	}

	return fmt.Errorf("%w: status=%d body=%s", ErrInvalidRequest, resp.StatusCode, string(body))
}

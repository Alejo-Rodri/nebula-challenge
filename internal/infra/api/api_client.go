package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func parseJSON(r io.Reader, v any) error {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	return json.NewDecoder(r).Decode(v)
}

func printError(verb string, endpoint string, errorType error, errorThrown error) error {
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

func (c *ApiClient) Info() (ApiInfoResponse, error) {
	result, err := get[ApiInfoResponse](c, "/info", nil)
	if err != nil {
		return result, err
	}

	return result, nil
}

// analyze func should be the basic flow
// first request -> retry until READY or some ERROR

func (c *ApiClient) Analyze(host string) (ApiAnalyzeResponse, error) {
	var endpoint string = "/analyze"

	baseURL, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return ApiAnalyzeResponse{}, printError("GET", endpoint, ErrParsingUrlRequest, err)
	}

	// first request
	// This parameter should be used only once to initiate a new assessment; further invocations should omit it to avoid causing an assessment loop.
	query := baseURL.Query()
	query.Set("host", host)
	query.Set("startNew", "on")
	query.Set("all", "done")

	result, err := get[ApiAnalyzeResponse](c, endpoint, query)
	if err != nil {
		return result, err
	}

	// TODO call analyze periodically
	// in which point do I start polling?
	// there should be another switch with the errors 
	
	// TODO randomize the delay
	// if 503 → 15min idle
	// if 529 → 30min idle
	// if 429 → stops everything
	query.Set("startNew", "off")
	for result.Status != "READY" {
		switch result.Status {
		case "ERROR":
			return result, fmt.Errorf("%w: status=%s, body=%+v", ErrInvalidResponse, result.Status, result)
		case "DNS":
			// poll cada 5s
			fmt.Println("sleeping 5s for DNS")
			time.Sleep(5 * time.Second)
			result, err := get[ApiAnalyzeResponse](c, endpoint, query)
			if err != nil {
				return result, err
			}
		case "IN_PROGRESS":
			// poll cada 15s
			fmt.Println("sleeping 15s for IN_PROGRESS")
			time.Sleep(15 * time.Second)
			result, err := get[ApiAnalyzeResponse](c, endpoint, query)
			if err != nil {
				return result, err
			}
		}
	}

	return result, nil
}

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

	resp, err := c.http.Get(baseURL.String())
	if err != nil {
		return result, printError("GET", endpoint, ErrConnection, err)
	}

	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return result, err
	}

	if err := parseJSON(resp.Body, &result); err != nil {
		return result, printError("GET", endpoint, ErrInvalidResponse, err)
	}

	return result, nil
}

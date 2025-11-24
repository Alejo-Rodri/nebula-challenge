package api

import (
	"fmt"
	"net/url"
	"time"
)

func (c *ApiClient) Info() (ApiInfoResponse, error) {
	result, err := get[ApiInfoResponse](c, "/info", nil)
	if err != nil {
		return result, err
	}

	return result, nil
}

// basic flow
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

		case "IN_PROGRESS":
			// poll cada 15s
			fmt.Println("sleeping 15s for IN_PROGRESS")
			time.Sleep(15 * time.Second)

		default:
			fmt.Printf("unknown status %s, sleeping 5s\n", result.Status)
        	time.Sleep(5 * time.Second)
		}

		result, err = get[ApiAnalyzeResponse](c, endpoint, query)
		fmt.Printf("status raw: %q\n", result.Status)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

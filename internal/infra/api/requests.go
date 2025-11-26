package api

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

type GetAbstractRequest[T any] func(c *ApiClient, endpoint string, query url.Values) (T, error)

func (c *ApiClient) Info(get GetAbstractRequest[ApiInfoResponse]) (ApiInfoResponse, error) {
	result, err := get(c, "/info", nil)
	if err != nil {
		return result, err
	}

	return result, nil
}

// basic flow
// first request -> retry until READY or some ERROR
func (c *ApiClient) Analyze(
	host string,
	get GetAbstractRequest[ApiAnalyzeResponse],
	) (ApiAnalyzeResponse, error) {

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

	result, err := get(c, endpoint, query)
	if err != nil {
		return result, err
	}
	
	query.Set("startNew", "off")

	for result.Status != "READY" {
		if err := handleStatus(result.Status); err != nil {
			return result, err
		}

		result, err = get(c, endpoint, query)
		if err != nil {
			if backoffErr := backoff(err); backoffErr != nil {
				return result, backoffErr
			}
		}
	}

	// se le inyectaria la funcion para almacenar el resultado y aca se llamaria

	return result, nil
}

func handleStatus(status string) error {
    switch status {
	case "ERROR":
		return fmt.Errorf("%w: status=%s", ErrInvalidRequest, status)
    case "DNS":
        fmt.Println("sleeping 5s for DNS")
        time.Sleep(5 * time.Second)
    case "IN_PROGRESS":
        fmt.Println("sleeping 15s for IN_PROGRESS")
        time.Sleep(15 * time.Second)
    default:
        fmt.Printf("unknown status %s, sleeping 5s\n", status)
        time.Sleep(5 * time.Second)
    }

	return nil
}

func backoff(err error) error {
    switch {
	// if 503 → 15min idle
    case errors.Is(err, ErrNoAvailableService):
        fmt.Println("the service in SSL labs is not available, sleeping 15min")
        time.Sleep(15 * time.Minute)
        return nil

	// if 529 → 30min idle
    case errors.Is(err, ErrOverloadedService):
        fmt.Println("the service in SSL labs is overloaded right now, sleeping 30min")
        time.Sleep(30 * time.Minute)
        return nil
    }

    // if 429 or error → stops everything 
    return err
}

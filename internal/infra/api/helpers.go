package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
)

// c *ApiClient, 
func Get[T any](
		c *http.Client,
		base,
		endpoint string,
		query url.Values,
	) (T, error) {
	var result T

	baseURL, err := url.Parse(base + endpoint)
	if err != nil {
		return result, printError("GET", endpoint, ErrParsingUrlRequest, err)
	}

	if query == nil {
		baseURL.RawQuery = ""
	} else {
		baseURL.RawQuery = query.Encode()
	}

	fmt.Printf("url: %s\n", baseURL.String())
	resp, err := c.Get(baseURL.String())
	if err != nil {
		return result, printError("GET", endpoint, ErrConnection, err)
	}

	defer resp.Body.Close()

	//fmt.Printf("status code: %d\n", resp.StatusCode)
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
		return fmt.Errorf("%w: can't read body: %v", ErrReadingBodyResponse, err)
	}

	var errType error
	switch resp.StatusCode {
	case 400:
		errType = ErrInvocationError
	case 429:
		errType = ErrRequestRateTooHigh
	case 500:
		errType = ErrInternalApiError
	case 503:
		errType = ErrNoAvailableService
	case 529:
		errType = ErrOverloadedService
	default:
		errType = ErrInvalidRequest
	}

	var apiErr ApiErrorsResponse
	if err := json.Unmarshal(body, &apiErr); err == nil {
		return fmt.Errorf("%w: %v", errType, apiErr)
	}

	return fmt.Errorf("%w: status=%d body=%s", errType, resp.StatusCode, string(body))
}

func mapInfo(r ApiInfoResponse) app.Info {
    return app.Info{
        EngineVersion:        r.EngineVersion,
        CriteriaVersion:      r.CriteriaVersion,
        ClientMaxAssessments: r.ClientMaxAssessments,
        MaxAssessments:       r.MaxAssessments,
        CurrentAssessments:   r.CurrentAssessments,
        NewAssessmentCoolOff: r.NewAssessmentCoolOff,
        Messages:             r.Messages,
    }
}

func mapAnalysis(r ApiAnalyzeResponse) app.Analysis {
	return app.Analysis{}
}
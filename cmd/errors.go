package cmd

import (
	"errors"
	"fmt"

	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
)

func HumanizeError(err error) string {
	switch {
	case errors.Is(err, api.ErrConnection):
		return "can't connect to the API"

	case errors.Is(err, api.ErrCreatingRequest):
		return "failed to create HTTP request"

	case errors.Is(err, api.ErrParsingUrlRequest):
		return "invalid request URL"

	case errors.Is(err, api.ErrReadingBodyResponse):
		return "failed to read server response"

	case errors.Is(err, api.ErrTimeout):
		return "API timeout"

	case errors.Is(err, api.ErrInvalidResponse):
		return "server returned an invalid response"

	case errors.Is(err, api.ErrInvalidRequest):
		return "client sent an invalid request"

	// API status code errors
	case errors.Is(err, api.ErrInvocationError):
		return "API rejected the request (400 - invocation error)"

	case errors.Is(err, api.ErrRequestRateTooHigh):
		return "rate limit exceeded (429 - too many requests)"

	case errors.Is(err, api.ErrInternalApiError):
		return "server crashed (500 - internal API error)"

	case errors.Is(err, api.ErrNoAvailableService):
		return "API is unavailable (503 - service unavailable)"

	case errors.Is(err, api.ErrOverloadedService):
		return "API is overloaded (529 - overloaded)"

	default:
		return fmt.Sprintf("unexpected error: %v", err)
	}
}

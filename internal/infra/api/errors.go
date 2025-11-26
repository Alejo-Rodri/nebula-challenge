package api

import "errors"

var (
	ErrConnection = errors.New("api connection failure")
	ErrCreatingRequest = errors.New("couldn't create a new request")
	ErrParsingUrlRequest = errors.New("couldn't parse correctly the url of the request")
	ErrReadingBodyResponse = errors.New("couldn't read the body of the response")
	ErrTimeout = errors.New("api timeout")
	ErrInvalidResponse = errors.New("invalid api response")
	ErrInvalidRequest = errors.New("invalid api request")

	// api error response status code
	ErrInvocationError = errors.New("400 - invocation error")
	ErrRequestRateTooHigh = errors.New("429 - client request rate too high or too many new assessments too fast")
	ErrInternalApiError = errors.New("500 - internal error")
	ErrNoAvailableService = errors.New("503 - the service is not available")
	ErrOverloadedService = errors.New("529 - the service is overloaded")
)

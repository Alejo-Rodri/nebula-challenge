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
)

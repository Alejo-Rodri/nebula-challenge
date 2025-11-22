package api

import "errors"

var (
	ErrConnection = errors.New("api connection failure")
	ErrTimeout = errors.New("api timeout")
	ErrInvalidResponse = errors.New("invalid api response")
	ErrInvalidRequest = errors.New("invalid api request")
)
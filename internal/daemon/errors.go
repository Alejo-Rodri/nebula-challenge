package daemon

import "errors"

var (
	ErrListRequest = errors.New("error requesting the assessments to the daemon")
	ErrMarshaling = errors.New("error marshaling struct")
	ErrWritingBody = errors.New("error writing payload to response")
)
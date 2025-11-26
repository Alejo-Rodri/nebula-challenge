package daemon

import "errors"

var (
	ErrListRequest = errors.New("error requesting the assessments to the daemon")
)
package db

import "errors"

var (
	ErrNotFound = errors.New("the key does not exist")
)
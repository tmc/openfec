package openfec

import "errors"

var (
	ErrUnauthorized = errors.New("openfec: unauthorized")
	ErrNotFound     = errors.New("openfec: not found")
)

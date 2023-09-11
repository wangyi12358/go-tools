package copy

import "errors"

var (
	ErrUnaddressable = errors.New("copy target is unaddressable")
	ErrInvalidSource = errors.New("copy source is invalid")
)

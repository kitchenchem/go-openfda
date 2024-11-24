package fda

import (
	"errors"
)

var ErrNotFound = errors.New("404 not found")
var ErrParameter = errors.New("fda error parameter")
var ErrTooManyRequests = errors.New("too many requests")

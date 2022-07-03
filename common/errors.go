package common

import "errors"

var (
	NotAuthorized = errors.New("not authorized")
	NotFound      = errors.New("not found")
)

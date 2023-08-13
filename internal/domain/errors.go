package domain

import "errors"

var (
	ErrorRequired = errors.New("required value")
	ErrorNotFound = errors.New("not found")
	ErrorNil      = errors.New("nil data")
)

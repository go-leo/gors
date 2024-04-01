package parser

import "errors"

var (
	ErrMultipleBodyBinding = errors.New("there are multiple body binding")

	ErrMultipleHttpMethod = errors.New("there are multiple methods")
)

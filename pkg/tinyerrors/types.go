package tinyerrors

import "errors"

type codeContainsError struct {
	Err  error
	code int
}

// Error to string converter...
func (e codeContainsError) Error() string {
	return e.Err.Error()
}

// Unwrap returns previous error...
func (e codeContainsError) Unwrap() error {
	return errors.Unwrap(e.Err)
}

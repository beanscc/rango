package ecode

import (
	"fmt"
)

// Error records an error contains code and message and detailed error that caused it
type Error struct {
	code    int    // code
	message string // msg
	detail  error  // detail err
}

func newError(code int, msg string, detail error) *Error {
	return &Error{code: code, message: msg, detail: detail}
}

// NewError return a new Error
func NewError(code int, msg string) *Error {
	return newError(code, msg, nil)
}

// NewErrorf return a new Error with a formatted message
func NewErrorf(code int, format string, args ...interface{}) *Error {
	return NewError(code, fmt.Sprintf(format, args...))
}

// NewErrorWithDetail return a new Error with detailed error that caused it
func NewErrorWithDetail(code int, msg string, detail error) *Error {
	return newError(code, msg, detail)
}

// FromError get Error from any error
func FromError(err error) (*Error, bool) {
	if err == nil {
		return nil, true
	}
	e, ok := err.(*Error)
	if ok {
		return e, true
	}
	return NewError(CodeUnknown, err.Error()), false
}

// Error implement error interface{}
func (e *Error) Error() string {
	if e.detail == nil {
		return fmt.Sprintf("error code: %d, msg: %s", e.code, e.message)
	}

	return fmt.Sprintf("error code: %d, msg: %s, detail: %s", e.code, e.message, e.detail)
}

// Code return error code
func (e *Error) Code() int {
	return e.code
}

// Msg return error msg
func (e *Error) Msg() string {
	return e.message
}

// Detail return detailed error
func (e *Error) Detail() error {
	return e.detail
}

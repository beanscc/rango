package ecode

import (
	"errors"
	"fmt"
)

// Error records an error contains code and message and detailed error that caused it
type Error struct {
	// 错误码，给开发者用的
	code int
	// 错误消息，给开发者用的
	message string // msg
	// 内部具体错误
	errors []error // detail err

	// 给用户展示的提示
}

// Code return error code
func (e *Error) Code() int {
	return e.code
}

// Message return error message
func (e *Error) Message() string {
	return e.message
}

// Error implement error interface{}
func (e *Error) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.code, e.message)
}

// Unwrap return inner details error
func (e *Error) Unwrap() []error {
	return e.errors
}

// Details return detailed error
func (e *Error) Details() []error {
	return e.errors
}

// SetMessage 重置 message
func (e *Error) SetMessage(msg string) *Error {
	e.message = msg
	return e
}

// WithDetails return a new Error with details error
func (e *Error) WithDetails(details ...error) *Error {
	e1 := &Error{code: e.code, message: e.message}
	e1.errors = append(e1.errors, details...)
	return e1
}

// New return a new Error
func New(code int, msg string) *Error {
	return &Error{code: code, message: msg}
}

// Newf return a new Error with a formatted message
func Newf(code int, format string, a ...any) *Error {
	return &Error{code: code, message: fmt.Sprintf(format, a...)}
}

// func Code(err error) int {
// 	e, ok := FromError(err)
// 	if ok {
// 		return e.Code()
// 	}
//
// 	return CodeUnknown
// }
//
// func FromCode(code int) *Error {
// 	return &Error{code: code}
// }

// FromError get Error from any error
func FromError(err error) (e *Error, ok bool) {
	if err == nil {
		return nil, true
	}
	ok = errors.As(err, &e)
	if ok {
		return e, true
	}
	return New(CodeUnknown, err.Error()), false
}

// Package errors deals with all types of error encountered.
package errors

import (
	"fmt"
)

// CodeType represents the type of the error.
type CodeType int

// IsOK returns true if the code type is OK.
func (code CodeType) IsOK() bool {
	return code == CodeOK
}

// Error interface for all errors
type Error interface {
	Error() string
	CodeType() CodeType
}

// NewError creates a new Error
func NewError(code CodeType, msg string) Error {
	return newError(code, msg)
}

type serverError struct {
	code CodeType
	msg  string
}

func newError(code CodeType, msg string) *serverError {
	if msg == "" {
		msg = CodeToDefaultMsg(code)
	}
	return &serverError{
		code: code,
		msg:  msg,
	}
}

// Error returns error details.
func (err *serverError) Error() string {
	return fmt.Sprintf("%d:%s", err.code, err.msg)
}

// CodeType returns the code of error.
func (err *serverError) CodeType() CodeType {
	return err.code
}

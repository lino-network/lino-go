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
	AddBlockChainCode(bcCode uint32) Error
	AddBlockChainLog(bcLog string) Error
	BlockChainCode() uint32
	BlockChainLog() string
	AddCause(cause error) Error
	Cause() error
}

// NewError creates a new Error
func NewError(code CodeType, msg string) Error {
	return newError(code, msg)
}

type serverError struct {
	code           CodeType
	msg            string
	blockChainCode uint32
	blockChainLog  string
	cause          error
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

// BlockChainCode returns the code of blockchain error.
func (err *serverError) AddBlockChainCode(bcCode uint32) Error {
	err.blockChainCode = bcCode
	return err
}

// BlockChainLog returns the log of blockchain error.
func (err *serverError) AddBlockChainLog(bcLog string) Error {
	err.blockChainLog = bcLog
	return err
}

// BlockChainCode returns the code of blockchain error.
func (err *serverError) BlockChainCode() uint32 {
	return err.blockChainCode
}

// BlockChainLog returns the log of blockchain error.
func (err *serverError) BlockChainLog() string {
	return err.blockChainLog
}

// TraceCause adds cause error.
func (err *serverError) AddCause(cause error) Error {
	err.cause = cause
	return err
}

// Cause returns the cause of error.
func (err *serverError) Cause() error {
	return err.cause
}

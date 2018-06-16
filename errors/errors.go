// Package errors deals with all types of error encountered in the backend.
package errors

import (
	"fmt"
	"runtime"
)

// CodeType represents the type of the error.
type CodeType int

// IsOK returns true if the code type is OK.
func (code CodeType) IsOK() bool {
	return code == CodeOK
}

// Error interface for all DLive errors
type Error interface {
	Error() string
	VerboseError() string
	CodeType() CodeType
	Trace(msg string) Error
	Tracef(msg string, args ...interface{}) Error
	TraceCause(cause error, msg string) Error
	Cause() error
}

// NewError creates a new Error
func NewError(code CodeType, msg string) Error {
	return newError(code, msg)
}

type traceItem struct {
	msg      string
	filename string
	lineno   int
}

func (ti traceItem) String() string {
	return fmt.Sprintf("%v:%v %v", ti.filename, ti.lineno, ti.msg)
}

// serverError is the customized Error used in the backend.
type serverError struct {
	code   CodeType
	msg    string
	cause  error
	traces []traceItem
}

func newError(code CodeType, msg string) *serverError {
	// TODO capture stacktrace if ENV is set.
	if msg == "" {
		msg = CodeToDefaultMsg(code)
	}
	return &serverError{
		code:   code,
		msg:    msg,
		cause:  nil,
		traces: nil,
	}
}

// Error returns error details.
func (err *serverError) Error() string {
	return fmt.Sprintf("%d:%s", err.code, err.msg)
}

// VerboseError returns all error stacks and traces
func (err *serverError) VerboseError() string {
	traceLog := ""
	for _, ti := range err.traces {
		traceLog += ti.String() + "\n"
	}
	return fmt.Sprintf("Error{%d:%s\nCause:%+v\ntrace:\n%v}", err.code, err.msg, err.cause, traceLog)
}

// CodeType returns the code of error.
func (err *serverError) CodeType() CodeType {
	return err.code
}

// Trace adds tracing information with msg.
func (err *serverError) Trace(msg string) Error {
	return err.doTrace(msg, 2)
}

// Tracef adds tracing information with formatted msg.
func (err *serverError) Tracef(format string, arg ...interface{}) Error {
	msg := fmt.Sprintf(format, arg...)
	return err.doTrace(msg, 2)
}

// TraceCause adds tracing information with cause and msg.
func (err *serverError) TraceCause(cause error, msg string) Error {
	err.cause = cause
	return err.doTrace(msg, 2)
}

func (err *serverError) doTrace(msg string, depth int) Error {
	_, fileName, line, ok := runtime.Caller(depth)
	if !ok {
		if fileName == "" {
			fileName = "<unknown>"
		}
		if line <= 0 {
			line = -1
		}
	}
	// Do not include the whole stack trace.
	err.traces = append(err.traces, traceItem{
		filename: fileName,
		lineno:   line,
		msg:      msg,
	})
	return err
}

// Cause returns the cause of error.
func (err *serverError) Cause() error {
	return err.cause
}

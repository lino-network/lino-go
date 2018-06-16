package errors

import "fmt"

// NOTE: Don't stringer this, we'll put better messages in later.
func CodeToDefaultMsg(code CodeType) string {
	switch code {
	case CodeFailedToBroadcast:
		return "Failed To Broadcast"
	case CodeInvalidSeqNumber:
		return "Invalid Seq Number"
	case CodeCheckTxFail:
		return "Check Tx Error"
	case CodeDeliverTxFail:
		return "Deliver Tx Error"
	default:
		return fmt.Sprintf("Unknown code %d", code)
	}
}

//FailedToBroadcast creates an error with CodeFailedToBroadcast
func FailedToBroadcast(msg string) Error {
	return newError(CodeFailedToBroadcast, msg)
}

//FailedToBroadcastf creates an error with CodeFailedToBroadcast and formatted message
func FailedToBroadcastf(format string, args ...interface{}) Error {
	return newError(CodeFailedToBroadcast, fmt.Sprintf(format, args...))
}

//InvalidSeqNumber creates an error with CodeInvalidSeqNumber
func InvalidSeqNumber(msg string) Error {
	return newError(CodeInvalidSeqNumber, msg)
}

//InvalidSeqNumberf creates an error with CodeInvalidSeqNumber and formatted message
func InvalidSeqNumberf(format string, args ...interface{}) Error {
	return newError(CodeInvalidSeqNumber, fmt.Sprintf(format, args...))
}

//CheckTxFail creates an error with CodeCheckTxFail
func CheckTxFail(msg string) Error {
	return newError(CodeCheckTxFail, msg)
}

//CheckTxFailf creates an error with CodeCheckTxFail and formatted message
func CheckTxFailf(format string, args ...interface{}) Error {
	return newError(CodeCheckTxFail, fmt.Sprintf(format, args...))
}

//DeliverTxFail creates an error with CodeDeliverTxFail
func DeliverTxFail(msg string) Error {
	return newError(CodeDeliverTxFail, msg)
}

//DeliverTxFailf creates an error with CodeDeliverTxFail and formatted message
func DeliverTxFailf(format string, args ...interface{}) Error {
	return newError(CodeDeliverTxFail, fmt.Sprintf(format, args...))
}

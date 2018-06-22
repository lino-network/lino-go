package errors

import "fmt"

func CodeToDefaultMsg(code CodeType) string {
	switch code {
	case CodeQueryFail:
		return "Failed To Query"
	case CodeFailedToBroadcast:
		return "Failed To Broadcast"
	case CodeCheckTxFail:
		return "Check Tx Error"
	case CodeDeliverTxFail:
		return "Deliver Tx Error"
	case CodeFialedToGetPubKeyFromHex:
		return "Failed To Get Pub Key From Hex"
	case CodeFailedToGetPrivKeyFromHex:
		return "Failed To Get Priv Key From Hex"
	case CodeInvalidArg:
		return "Invalid argument"
	case CodeEmptyResponse:
		return "Empty Response"
	default:
		return fmt.Sprintf("Unknown code %d", code)
	}
}

//CodeQueryFail creates an error with CodeQueryFail
func QueryFail(msg string) Error {
	return newError(CodeQueryFail, msg)
}

//QueryFailf creates an error with CodeQueryFail and formatted message
func QueryFailf(format string, args ...interface{}) Error {
	return newError(CodeQueryFail, fmt.Sprintf(format, args...))
}

//FailedToBroadcast creates an error with CodeFailedToBroadcast
func FailedToBroadcast(msg string) Error {
	return newError(CodeFailedToBroadcast, msg)
}

//FailedToBroadcastf creates an error with CodeFailedToBroadcast and formatted message
func FailedToBroadcastf(format string, args ...interface{}) Error {
	return newError(CodeFailedToBroadcast, fmt.Sprintf(format, args...))
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

//FailedToGetPubKeyFromHex creates an error with CodeFialedToGetPubKeyFromHex
func FailedToGetPubKeyFromHex(msg string) Error {
	return newError(CodeFialedToGetPubKeyFromHex, msg)
}

//FailedToGetPubKeyFromHex creates an error with CodeDeliverTxFail and formatted message
func FailedToGetPubKeyFromHexf(format string, args ...interface{}) Error {
	return newError(CodeFialedToGetPubKeyFromHex, fmt.Sprintf(format, args...))
}

//FailedToGetPrivKeyFromHex creates an error with CodeFailedToGetPrivKeyFromHex
func FailedToGetPrivKeyFromHex(msg string) Error {
	return newError(CodeFailedToGetPrivKeyFromHex, msg)
}

//FailedToGetPrivKeyFromHexf creates an error with CodeFailedToGetPrivKeyFromHex and formatted message
func FailedToGetPrivKeyFromHexf(format string, args ...interface{}) Error {
	return newError(CodeFailedToGetPrivKeyFromHex, fmt.Sprintf(format, args...))
}

//InvalidArg creates an error with CodeInvalidArg
func InvalidArg(msg string) Error {
	return newError(CodeInvalidArg, msg)
}

//InvalidArgf creates an error with CodeInvalidArg and formatted message
func InvalidArgf(format string, args ...interface{}) Error {
	return newError(CodeInvalidArg, fmt.Sprintf(format, args...))
}

//EmptyResponse creates an error with CodeEmptyResponse
func EmptyResponse(msg string) Error {
	return newError(CodeEmptyResponse, msg)
}

//EmptyResponsef creates an error with CodeEmptyResponse and formatted message
func EmptyResponsef(format string, args ...interface{}) Error {
	return newError(CodeEmptyResponse, fmt.Sprintf(format, args...))
}

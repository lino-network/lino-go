package errors

// Code types.
const (
	CodeOK CodeType = iota // 0
	CodeQueryFail
	CodeFailedToBroadcast
	CodeCheckTxFail
	CodeDeliverTxFail
	CodeFailedToGetPubKeyFromHex // 5
	CodeFailedToGetPrivKeyFromHex
	CodeInvalidArg
	CodeInvalidNodeURL
	CodeInvalidSequenceNumber
	CodeEmptyResponse // 10
)

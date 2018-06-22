package errors

// Code types.
const (
	CodeOK CodeType = iota // 0
	CodeFailedToBroadcast
	CodeInvalidSeqNumber
	CodeCheckTxFail
	CodeDeliverTxFail
	CodeFialedToGetPubKeyFromHex
	CodeQueryFail
	CodeUnmarshalFail
	CodeFailedToGetPrivKeyFromHex
	CodeInvalidArg
)

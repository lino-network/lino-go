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
	CodeEmptyResponse    // 10
	CodeTimeout          // time out for one tx
	CodeBroadcastTimeout // time out for each single broadcast
	CodeInvalidSignature // invalid sign bytes
)

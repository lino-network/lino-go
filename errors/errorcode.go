package errors

// Code types.
const (
	CodeOK CodeType = 0 // 0

	CodeFailedToBroadcast = 1
	CodeInvalidSeqNumber  = 2
	CodeCheckTxFail       = 3
	CodeDeliverTxFail     = 4
)

package model

import (
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	"time"
)

//
// block related
//

// type Block struct {
// 	Header     tmtypes.Header       `json:"header"`
// 	Evidence   tmtypes.EvidenceData `json:"evidence"`
// 	LastCommit *tmtypes.Commit      `json:"last_commit"`
// 	Data       *Data                `json:"data"`
// }

type BlockStatus struct {
	LatestBlockHeight int64     `json:"latest_block_height"`
	LatestBlockTime   time.Time `json:"latest_block_time"`
}

// type Data struct {
// 	Txs Txs `json:"txs"`
// }

// type Txs []Transaction

type BlockTx struct {
	Height int64      `json:"height"`
	Tx     auth.StdTx `json:"tx"`
	Code   uint32     `json:"code"`
	Log    string     `json:"log"`
}

type BroadcastResponse struct {
	CommitHash string `json:"commit_hash"`
	Height     int64  `json:"height"`
}

type TxAndSequenceNumber struct {
	Username string             `json:"username"`
	Sequence uint64             `json:"sequence"`
	Tx       *TransactionResult `json:"tx"`
}

type TransactionResult struct {
	Hash   string `json:"hash"`
	Height int64  `json:"height"`
	Code   uint32 `json:"code"`
}

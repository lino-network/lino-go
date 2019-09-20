# Documentation
* [Init](#init)  
* [API](#api) 
    * [Node](#node)
        * [Get Lastest Block Height](#get-lastest-block-height)
        * [Get Lastest Block Time](#get-lastest-block-time)
        * [Check If Node is Synced with Latest Blocks](#check-if-node-is-synced-with-latest-blocks)
        * [Get Block Information](#get-block-information)
        * [Get All Transactions in a Block](#get-all-transactions-in-a-block)
        * [Get Transactions by Hash](#get-transactions-by-hash)
    * [Account](#account)
        * [Generate Private Key Pair](#generate-private-key-pair)
    * [Query](#query)
        * [Developer](#developer)  
        * [Infra](#infra)  
        * [Blockchain Parameters](#blockchain-parameters)  
        * [Post](#post)  
        * [Proposal](#proposal)  
        * [Block](#block)  
        * [Validator](#validator)  
        * [Vote](#vote)  
    * [Broadcast](#broadcast) 
        * [Synchronizing and Analyzing the Successful Transfers](#synchronizing-and-analyzing-the-successful-transfers)
        * [Account](#broadcast-account)  
        * [Post](#broadcast-post)  
        * [Validator](#broadcast-validator)  
        * [Vote](#broadcast-vote)  
        * [Developer](#broadcast-developer)  
        * [Infra](#broadcast-infra)  
        * [Proposal](#broadcast-proposal)  

## Init

To connect with latest Lino Blockchain, chain id and node url should be set specifically as following:

```
api := api.NewLinoAPIFromArgs(&api.Options{
	ChainID: "lino-testnet-upgrade2",
	NodeURL: "https://fullnode.lino.network:443",
})
```

## API
### Node
#### Get Lastest Block Height
```
    	resp, _ := api.GetBlockStatus(context.Background())
	fmt.Println(resp.SyncInfo.LatestBlockHeight)
```

#### Get Lastest Block Time
```
    	resp, _ := api.GetBlockStatus(context.Background())
	fmt.Println(resp.SyncInfo.LatestBlockTime)
```
#### Check If Node is Synced with Latest Blocks
```
    	resp, _ := api.GetBlockStatus(context.Background())
	fmt.Println(resp.SyncInfo.CatchingUp)
```

#### Get Block Information
```
    	resp, _ := api.GetBlock(context.Background(), 1)
```

Block Information include BlockID, all precommits, height, header, all transactions, etc.

#### Get All Transactions in a Block

```
package main

import (
	"context"
	"fmt"

	"github.com/lino-network/lino-go/api"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	linoapp "github.com/lino-network/lino/app"
)

func main() {
	api := api.NewLinoAPIFromArgs(&api.Options{
		ChainID: "lino-testnet-upgrade2",
		NodeURL: "https://fullnode.lino.network:443",
	})

	resp, _ := api.GetBlock(context.Background(), 24208)
    for i := 0, i < len(resp.Data.Txs); i ++ {
        var tx auth.StdTx
        cdc := linoapp.MakeCodec()
        if err := cdc.UnmarshalJSON(resp.Data.Txs[0], &tx); err != nil {
            panic(err)
        }
        fmt.Println(tx)
    }
}
```

#### Get Transactions by Hash

```
	commitHash, _ := hex.DecodeString("df6bf5c9cfc8b2a999dcd6544218f972a557faab43439ff81047041cb980ec59")
	tx, _ := api.GetTx(context.Background(), commitHash)
	fmt.Println(tx.Code)   // error code of tx execution result. 0 means success.
	fmt.Println(tx.Height) // height when tx is commited.
	fmt.Println(tx.Log)    // log of tx execution result.
	fmt.Println(tx.Tx)     // raw transaction
```

### Account

#### Generate Private Key Pair
```
package main

import (
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func main() {
	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()
	addr := pub.Address()
}
```

#### Get AccountInfo By Username
```
accountInfo, err := api.GetAccountInfo(ctx, username)
```

#### Get Transaction Public Key
```
txPubKey, err := api.GetTransactionPubKey(ctx, username)
```

#### Get Signing Public Key
```
signingPubKey, err := api.GetSigningPubKey(ctx, username)
```

#### Check Does Username Match Transaction Private Key
```
isMatch, err := api.DoesUsernameMatchTxPrivKey(ctx, username, txPrivKeyHex)
```

#### Check Does Username Match Signing Private Key
```
isMatch, err := api.DoesUsernameMatchAppPrivKey(ctx, username, appPrivKeyHex)
```

#### Get AccountBank By Username
```
accountBank, err := api.GetAccountBank(ctx, username)
```

#### Get AccountBank By Address
```
accountBank, err := api.GetAccountBankByAddress(ctx, address)
```

#### Get AccountMeta
```
accountMeta, err := api.GetAccountMeta(ctx, username)
```

#### Get Next Sequence Number
```
seq, err := api.GetSeqNumber(ctx, username)
```

#### Get Granted Public Key
```
grantPubKey, err := api.GetGrantPubKey(ctx, username, grantTo, permission)
```

#### Get All Granted Public Keys
```
pubKeyToGrantPubKeyMap, err := api.GetAllGrantPubKeys(ctx, username)
```

### Developer
#### Get Developer 
```
developer, err := api.GetDeveloper(ctx, developerName)
```

### Infra
#### Get Infra Provider
```
infraProvider, err := api.GetInfraProvider(ctx, providerName)
```

### Blockchain Parameters
#### Get Evaluate Of Content Value Param
```
p, err := api.GetEvaluateOfContentValueParam(ctx)
```
#### Get Global Allocation Param
```
p, err := api.GetGlobalAllocationParam(ctx)
```
#### Get Infra Internal Allocation Param
```
p, err := api.GetInfraInternalAllocation(ctx)
```
#### Get Developer Param
```
p, err := api.GetDeveloperParam(ctx)
```
#### Get Vote Param
```
p, err := api.GetVoteParam(ctx)
```
#### Get Proposal Param
```
p, err := api.GetProposalParam(ctx)
```
#### Get Validator Param
```
p, err := api.GetValidatorParam(ctx)
```
#### Get Bandwidth Param
```
p, err := api.GetBandwidthParam(ctx)
```
#### Get Account Param
```
p, err := api.GetAccountParam(ctx)
```
#### Get Post Param
```
p, err := api.GetPostParam(ctx)
```

### Post
#### Get PostInfo
```
postInfo, err := api.GetPostInfo(ctx, author, postID)
```
### Proposal
#### Get Proposal List
```
proposalList, err := api.GetProposalList(ctx)
```
#### Get Proposal 
```
proposal, err := api.GetProposal(ctx, proposalID)
```
#### Get Ongoing Proposals
```
ongoingProposals, err := api.GetOngoingProposal(ctx)
```
#### Get Expired Proposals
```
expiredProposals, err := api.GetExpiredProposal(ctx)
```

### Block
#### Get Block
```
block, err := api.GetBlock(ctx, height)
```
#### Get Block Status
```
blockStatus, err := api.GetBlockStatus(ctx)
```

### Validator
#### Get Validator
```
validator, err := api.GetValidator(ctx, username)
```
#### Get All Validators
```
validators, err := api.GetAllValidators(ctx)
```

### Vote
#### Get Voter
```
voter, err := api.GetVoter(ctx, voterName)
```

### Broadcast
### Synchronizing and Analyzing the Successful Transfers
```
resp, err := api.Transfer(ctx, sender, receiver, amount, memo, privKeyHex)
if err != nil {
    panic(err) // transfer failed
}
fmt.Println(resp.CommitHash) // commit hash of the transaction, can be queried by tx query api.
fmt.Println(resp.Height) // height when the transaction was executed
```

### Broadcast Account
#### Register A New User
```
resp, err := api.Register(ctx, referrer, registerFee, newUsername, newUserResetPubHex, newUserTxPubHex, newUserAppPubHex, referrerTxPrivKey)
```
#### Transfer LINO Between two users
```
resp, err := api.Transfer(ctx, sender, receiver, amount, memo, privKeyHex)
```
#### Follow 
```
resp, err := api.Follow(ctx, follower, followee, privKeyHex)
```
#### Unfollow 
```
resp, err := api.Unfollow(ctx, follower, followee, privKeyHex)
```
#### Claim Reward
```
resp, err := api.Claim(ctx, username, privKeyHex)
```
#### Claim Interest
```
resp, err := api.ClaimInterest(ctx, username, privKeyHex)
```
#### Update Account
```
resp, err := api.UpdateAccount(ctx, username, jsonMeta, privKeyHex)
```
#### Recover 
```
resp, err := api.Recover(ctx, username, newResetPubKeyHex, newTransactionPubKeyHex, newAppPubKeyHex, privKeyHex)
```

### Broadcast Post
#### Create Post
```
resp, err := api.CreatePost(ctx, author, postID, title, content, parentAuthor, parentPostID, sourceAuthor, sourcePostID, redistributionSplitRate, links, privKeyHex)
```
#### Donate To A Post
```
resp, err := api.Donate(ctx, username, author, amount, postID, fromApp, memo, privKeyHex)
```
#### ReportOrUpvote To A Post
```
resp, err := api.ReportOrUpvote(ctx, username, author, postID, isReport, privKeyHex)
```
#### Delete Post
```
resp, err := api.DeletePost(ctx, author, postID, privKeyHex)
```
#### View A Post
```
resp, err := api.View(ctx, username, author, postID, privKeyHex)
```
#### Update Post
```
resp, err := api.UpdatePost(ctx, author, title, postID, content, links, privKeyHex)
```

### Broadcast Validator
#### Validator Deposit
```
resp, err := api.ValidatorDeposit(ctx, username, deposit, validatorPubKey, link, privKeyHex)
```
#### Validator Withdraw
```
resp, err := api.ValidatorWithdraw(ctx, username, amount, privKeyHex)
```
#### Validator Revoke
```
resp, err := api.ValidatorRevoke(ctx, username, privKeyHex)
```

### Broadcast Vote
#### Voter StakeIn
```
resp, err := api.StakeIn(ctx, username, deposit, privKeyHex)
```
#### Voter StakeOut
```
resp, err := api.StakeOut(ctx, username, amount, privKeyHex)
```
#### Delegate To Voter
```
resp, err := api.Delegate(ctx, delegator, voter, amount, privKeyHex)
```
#### Delegator Withdraw
```
resp, err := api.DelegatorWithdraw(ctx, delegator, voter, amount, privKeyHex)
```
#### RevokeDelegation
```
resp, err := api.RevokeDelegation(ctx, delegator, voter, privKeyHex)
```

### Broadcast Developer
#### Developer Register
```
resp, err := api.DeveloperRegister(ctx, username, deposit, website, description, appMetaData, privKeyHex)
```
#### DeveloperUpdate
```
resp, err := api.DeveloperUpdate(ctx, username, website, description, appMetaData, privKeyHex)
```
#### DeveloperRevoke
```
resp, err := api.DeveloperRevoke(ctx, username, privKeyHex)
```
#### Grant Permission
```
resp, err := api.GrantPermission(ctx, username, authorizedApp, validityPeriodSec, grantLevel, privKeyHex)
```
#### Pre Authorization Permission
```
resp, err := api.PreAuthorizationPermission(ctx, username, authorizedApp, validityPeriodSec, amount, privKeyHex)
```
#### Revoke Permission
```
resp, err := api.RevokePermission(ctx, username, pubKeyHex, privKeyHex)
```

### Broadcast Infra
#### Infra Provider Report
```
resp, err := api.ProviderReport(ctx, username, usage, privKeyHex)
```

### Broadcast Proposal
#### Change Evaluate Of Content Value Param
```
resp, err := api.ChangeEvaluateOfContentValueParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Global Allocation Param
```
resp, err := api.ChangeGlovalAllocationParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Infra Internal Allocation Param
```
resp, err := api.ChangeInfraInternalAllocationParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Vote Param
```
resp, err := api.ChangeVoteParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Proposal Param
```
resp, err := api.ChangeProposalParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Developer Param
```
resp, err := api.ChangeDeveloperParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Validator Param
```
resp, err := api.ChangeValidatorParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Bandwidth Param
```
resp, err := api.ChangeBandwidthParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Account Param
```
resp, err := api.ChangeAccountParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Change Post Param
```
resp, err := api.ChangePostParam(ctx, creator, parameter, reason, privKeyHex)
```
#### Delete Post Content
```
resp, err := api.DeletePostContent(ctx, creator, postAuthor, postID, reason, privKeyHex)
```
#### Vote Proposal
```
resp, err := api.VoteProposal(ctx, voter, proposalID, result, privKeyHex)
```

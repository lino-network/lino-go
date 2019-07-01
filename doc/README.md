# Documentation
* [Get Tools & dependencies](#get-tools-dependencies)  
* [Init](#init)  
* [API](#api)  
    * [Query](#query)  
        * [Account](#account)  
        * [Developer](#developer)  
        * [Infra](#infra)  
        * [Blockchain Parameters](#blockchain-parameters)  
        * [Post](#post)  
        * [Proposal](#proposal)  
        * [Block](#block)  
        * [Validator](#validator)  
        * [Vote](#vote)  
    * [Broadcast](#broadcast)  
        * [Account](#broadcast-account)  
        * [Post](#broadcast-post)  
        * [Validator](#broadcast-validator)  
        * [Vote](#broadcast-vote)  
        * [Developer](#broadcast-developer)  
        * [Infra](#broadcast-infra)  
        * [Proposal](#broadcast-proposal)  

## Get Tools & Dependencies
```
dep ensure
```

## Init
```
api := api.NewLinoAPIFromArgs(chainID, nodeURL)
```
chanID and nodeURL can be found remotely from https://github.com/lino-network/testnets/blob/master/lino-testnet/genesis.json 
or locally from ~/.lino/config/genesis.json

For example,  
Remotely: chainID = "lino-testnet" and nodeURL = "https://fullnode.lino.network:443"  
Locally: chainID = "test-chain-q8lMWR" and nodeURL = "http://localhost:26657"  

## API

### Query
#### Account 
##### Get AccountInfo
```
accountInfo, err := api.GetAccountInfo(ctx, username)
```
##### Get Transaction Public Key
```
txPubKey, err := api.GetTransactionPubKey(ctx, username)
```
##### Get App Public Key
```
appPubKey, err := api.GetAppPubKey(ctx, username)
```
##### Check Does Username Match Reset Private Key
```
isMatch, err := api.DoesUsernameMatchResetPrivKey(ctx, username, resetPrivKeyHex)
```
##### Check Does Username Match Transaction Private Key
```
isMatch, err := api.DoesUsernameMatchTxPrivKey(ctx, username, txPrivKeyHex)
```
##### Check Does Username Match App Private Key
```
isMatch, err := api.DoesUsernameMatchAppPrivKey(ctx, username, appPrivKeyHex)
```
##### Get AccountBank
```
accountBank, err := api.GetAccountBank(ctx, username)
```
##### Get AccountMeta
```
accountMeta, err := api.GetAccountMeta(ctx, username)
```
##### Get Next Sequence Number
```
seq, err := api.GetSeqNumber(ctx, username)
```
##### Get All Balance History From All Buckets
```
allBalanceHistory, err := api.GetAllBalanceHistory(ctx, username)
```
##### Get A Certain Number Of Recent Balance History
```
recentBalanceHistory, err := api.GetRecentBalanceHistory(ctx, uesrname, numOfHistory)
```
##### Get Balance History In The Range Of Index [from, to] Inclusively
```
rangedBalanceHistory, err := api.GetBalanceHistoryFromTo(ctx, username, from, to)
```
##### Get Balance History From A Certain Bucket
```
bucketBalanceHistory, err := api.GetBalanceHistory(ctx, username, bucketIndex)
```
##### Get Granted Public Key
```
grantPubKey, err := api.GetGrantPubKey(ctx, username, pubKeyHex)
```
##### Get Reward
```
reward, err := api.GetReward(ctx, username)
```
##### Get Reward At A Certain Block Height
```
reward, err := api.GetRewardAtHeight(ctx, username, height)
```
##### Get All Reward History From All Buckets
```
allRewardHistory, err := api.GetAllRewardHistory(ctx, username)
```
##### Get A Certain Number Of Recent Reward History
```
recentRewardHistory, err := api.GetRecentRewardHistory(ctx, username, numOfHistory) 
```
##### Get Reward History In The Range Of Index [from, to] Inclusively
```
rangedRewardHistory, err := api.GetRewardHistoryFromTo(ctx, username, from, to)
```
##### Get Reward History From A Certain Bucket
```
bucketRewardHistory, err := api.GetRewardHistory(ctx, username, bucketIndex)
```
##### Get Donation Relationship
```
relationship, err := api.GetRelationship(ctx, me, other)
```
##### Get Follower Meta
```
followerMeta, err := api.GetFollowerMeta(ctx, me, myFollower)
```
##### Get Following Meta
```
followingMeta, err := api.GetFollowingMeta(ctx, me, myFollowing)
```
##### Get All Granted Public Keys
```
pubKeyToGrantPubKeyMap, err := api.GetAllGrantPubKeys(ctx, username)
```
##### Get All Donation Relationships 
```
userToRelationshipMap, err := api.GetAllRelationships(ctx, username)
```
##### Get All Follower Meta
```
followerToMetaMap, err := api.GetAllFollowerMeta(ctx, username)
```
##### Get All Following Meta
```
followingToMetaMap, err := api.GetAllFollowingMeta(ctx, username)
```

#### Developer
##### Get Developer 
```
developer, err := api.GetDeveloper(ctx, developerName)
```
##### Get All Developers
```
devevlopers, err := api.GetDevelopers(ctx)
```

#### Infra
##### Get Infra Provider
```
infraProvider, err := api.GetInfraProvider(ctx, providerName)
```
##### Get All Infra Providers
```
infraProviders, err := api.GetInfraProviders(ctx)
```

#### Blockchain Parameters
##### Get Evaluate Of Content Value Param
```
p, err := api.GetEvaluateOfContentValueParam(ctx)
```
##### Get Global Allocation Param
```
p, err := api.GetGlobalAllocationParam(ctx)
```
##### Get Infra Internal Allocation Param
```
p, err := api.GetInfraInternalAllocation(ctx)
```
##### Get Developer Param
```
p, err := api.GetDeveloperParam(ctx)
```
##### Get Vote Param
```
p, err := api.GetVoteParam(ctx)
```
##### Get Proposal Param
```
p, err := api.GetProposalParam(ctx)
```
##### Get Validator Param
```
p, err := api.GetValidatorParam(ctx)
```
##### Get Coin Day Param
```
p, err := api.GetCoinDayParam(ctx)
```
##### Get Bandwidth Param
```
p, err := api.GetBandwidthParam(ctx)
```
##### Get Account Param
```
p, err := api.GetAccountParam(ctx)
```
##### Get Post Param
```
p, err := api.GetPostParam(ctx)
```

#### Post
##### Get PostInfo
```
postInfo, err := api.GetPostInfo(ctx, author, postID)
```
##### Get PostMeta
```
postMeta, err := api.GetPostMeta(ctx, author, postID)
```
##### Get Post Comment
```
comment, err := api.GetPostComment(ctx, author, postID, commentPermlink)
```
##### Get Post View
```
view, err := api.GetPostView(ctx, author, postID, viewUser)
```
##### Get Post Donations
```
donations, err := api.GetPostDonations(ctx, author, postID, donateUser)
```
##### Get Post ReportOrUpvote
```
reportOrUpvote, err := api.GetPostReportOrUpvote(ctx, author, postID, user)
```
##### Get User All Posts
```
permlinkToPostMap, err := api.GetUserAllPosts(ctx, username)
```
##### Get Post All Comments
```
permlinkToCommentMap, err := api.GetPostAllComments(ctx, author, postID)
```
##### Get Post All Views
```
userToViewMap, err := api.GetPostAllViews(ctx, author, postID)
```
##### Get Post All Donations
```
userToDonationsMap, err := api.GetPostAllDonations(ctx, author, postID)
```
##### Get Post All ReportOrUpvotes
```
userToReportOrUpvoteMap, err := api.GetPostAllReportOrUpvotes(ctx, author, postID)
```

#### Proposal
##### Get Proposal List
```
proposalList, err := api.GetProposalList(ctx)
```
##### Get Proposal 
```
proposal, err := api.GetProposal(ctx, proposalID)
```
##### Get Ongoing Proposals
```
ongoingProposals, err := api.GetOngoingProposal(ctx)
```
##### Get Expired Proposals
```
expiredProposals, err := api.GetExpiredProposal(ctx)
```
##### Get Next Proposal ID
```
nextProposalID, err := api.GetNextProposalID(ctx)
```

#### Block
##### Get Block
```
block, err := api.GetBlock(ctx, height)
```
##### Get Block Status
```
blockStatus, err := api.GetBlockStatus(ctx)
```

#### Validator
##### Get Validator
```
validator, err := api.GetValidator(ctx, username)
```
##### Get All Validators
```
validators, err := api.GetAllValidators(ctx)
```

#### Vote
##### Get Delegation
```
delegation, err := api.GetDelegation(ctx, voter, delegator)
```
##### Get Voter All Delegations
```
delegations, err := api.GetVoterAllDelegation(ctx, voter)
```
##### Get Delegator All Delegations
```
delegations, err := api.GetDelegatorAllDelegation(ctx, delegatorName)
```
##### Get Voter
```
voter, err := api.GetVoter(ctx, voterName)
```
##### Get Vote
```
vote, err := api.GetVote(ctx, proposalID, voter)
```
##### Get Proposal All Votes
```
votes, err := api.GetProposalAllVotes(ctx, proposalID)
```

### Broadcast
#### Broadcast Account
##### Register A New User
```
seq, err := api.GetSeqNumber(ctx, referrer)
resp, err := api.Register(ctx, referrer, registerFee, newUsername, newUserResetPubHex, newUserTxPubHex, newUserAppPubHex, referrerTxPrivKey, seq)
```
##### Transfer LINO Between two users
```
seq, err := api.GetSeqNumber(ctx, sender)
resp, err := api.Transfer(ctx, sender, receiver, amount, memo, privKeyHex, seq)
```
##### Follow 
```
seq, err := api.GetSeqNumber(ctx, follower)
resp, err := api.Follow(ctx, follower, followee, privKeyHex, seq)
```
##### Unfollow 
```
seq, err := api.GetSeqNumber(ctx, follower)
resp, err := api.Unfollow(ctx, follower, followee, privKeyHex, seq)
```
##### Claim Reward
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.Claim(ctx, username, privKeyHex, seq)
```
##### Claim Interest
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.ClaimInterest(ctx, username, privKeyHex, seq)
```
##### Update Account
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.UpdateAccount(ctx, username, jsonMeta, privKeyHex, seq)
```
##### Recover 
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.Recover(ctx, username, newResetPubKeyHex, newTransactionPubKeyHex, newAppPubKeyHex, privKeyHex, seq)
```

#### Broadcast Post
##### Create Post
```
seq, err := api.GetSeqNumber(ctx, author)
resp, err := api.CreatePost(ctx, author, postID, title, content, parentAuthor, parentPostID, sourceAuthor, sourcePostID, redistributionSplitRate, links, privKeyHex, seq)
```
##### Donate To A Post
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.Donate(ctx, username, author, amount, postID, fromApp, memo, privKeyHex, seq)
```
##### ReportOrUpvote To A Post
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.ReportOrUpvote(ctx, username, author, postID, isReport, privKeyHex, seq)
```
##### Delete Post
```
seq, err := api.GetSeqNumber(ctx, author)
resp, err := api.DeletePost(ctx, author, postID, privKeyHex, seq)
```
##### View A Post
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.View(ctx, username, author, postID, privKeyHex, seq)
```
##### Update Post
```
seq, err := api.GetSeqNumber(ctx, author)
resp, err := api.UpdatePost(ctx, author, title, postID, content, links, privKeyHex, seq)
```

#### Broadcast Validator
##### Validator Deposit
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.ValidatorDeposit(ctx, username, deposit, validatorPubKey, link, privKeyHex, seq)
```
##### Validator Withdraw
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.ValidatorWithdraw(ctx, username, amount, privKeyHex, seq)
```
##### Validator Revoke
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.ValidatorRevoke(ctx, username, privKeyHex, seq)
```

#### Broadcast Vote
##### Voter StakeIn
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.StakeIn(ctx, username, deposit, privKeyHex, seq)
```
##### Voter StakeOut
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.StakeOut(ctx, username, amount, privKeyHex, seq)
```
##### Delegate To Voter
```
seq, err := api.GetSeqNumber(ctx, delegator)
resp, err := api.Delegate(ctx, delegator, voter, amount, privKeyHex, seq)
```
##### Delegator Withdraw
```
seq, err := api.GetSeqNumber(ctx, delegator)
resp, err := api.DelegatorWithdraw(ctx, delegator, voter, amount, privKeyHex, seq)
```
##### RevokeDelegation
```
seq, err := api.GetSeqNumber(ctx, delegator)
resp, err := api.RevokeDelegation(ctx, delegator, voter, privKeyHex, seq)
```

#### Broadcast Developer
##### Developer Register
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.DeveloperRegister(ctx, username, deposit, website, description, appMetaData, privKeyHex, seq)
```
##### DeveloperUpdate
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.DeveloperUpdate(ctx, username, website, description, appMetaData, privKeyHex, seq)
```
##### DeveloperRevoke
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.DeveloperRevoke(ctx, username, privKeyHex, seq)
```
##### Grant Permission
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.GrantPermission(ctx, username, authorizedApp, validityPeriodSec, grantLevel, privKeyHex, seq)
```
##### Pre Authorization Permission
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.PreAuthorizationPermission(ctx, username, authorizedApp, validityPeriodSec, amount, privKeyHex, seq)
```
##### Revoke Permission
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.RevokePermission(ctx, username, pubKeyHex, privKeyHex, seq)
```

#### Broadcast Infra
##### Infra Provider Report
```
seq, err := api.GetSeqNumber(ctx, username)
resp, err := api.ProviderReport(ctx, username, usage, privKeyHex, seq)
```

#### Broadcast Proposal
##### Change Evaluate Of Content Value Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeEvaluateOfContentValueParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Global Allocation Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeGlovalAllocationParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Infra Internal Allocation Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeInfraInternalAllocationParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Vote Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeVoteParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Proposal Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeProposalParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Developer Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeDeveloperParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Validator Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeValidatorParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Bandwidth Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeBandwidthParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Account Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangeAccountParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Change Post Param
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.ChangePostParam(ctx, creator, parameter, reason, privKeyHex, seq)
```
##### Delete Post Content
```
seq, err := api.GetSeqNumber(ctx, creator)
resp, err := api.DeletePostContent(ctx, creator, postAuthor, postID, reason, privKeyHex, seq)
```
##### Vote Proposal
```
seq, err := api.GetSeqNumber(ctx, voter)
resp, err := api.VoteProposal(ctx, voter, proposalID, result, privKeyHex, seq)
```

# Documentation
[Get Tools & dependencies](#get Tools & Dependencies)  
[Init](#init)  
[API](#aPI)  

## Get Tools & Dependencies
```
dep ensure
```

## Init
```
api := api.NewLinoAPIFromArgs(chainID, nodeURL)
```
chanID and nodeURL can be found remotely from https://linotracker.io/ 
or locally from ~/.lino/config/genesis.json

For example,  
Remotely: chainID = "test-chain-BgWrtq" and nodeURL = "http://fullnode.linovalidator.io:80"  
Locally: chainID = "test-chain-q8lMWR" and nodeURL = "http://localhost:26657"  

## API

### Query
#### Account 
##### Get AccountInfo
```
accountInfo, err := api.GetAccountInfo(username)
```
##### Get Transaction Public Key
```
txPubKey, err := api.GetTransactionPubKey(username)
```
##### Get App Public Key
```
appPubKey, err := api.GetAppPubKey(username)
```
##### Check Does Username Match Reset Private Key
```
isMatch, err := api.DoesUsernameMatchResetPrivKey(username, resetPrivKeyHex)
```
##### Check Does Username Match Transaction Private Key
```
isMatch, err := api.DoesUsernameMatchTxPrivKey(username, txPrivKeyHex)
```
##### Check Does Username Match App Private Key
```
isMatch, err := api.DoesUsernameMatchAppPrivKey(username, appPrivKeyHex)
```
##### Get AccountBank
```
accountBank, err := api.GetAccountBank(username)
```
##### Get AccountMeta
```
accountMeta, err := api.GetAccountMeta(username)
```
##### Get Next Sequence Number
```
seq, err := api.GetSeqNumber(username)
```
##### Get All Balance History From All Buckets
```
allBalanceHistory, err := api.GetAllBalanceHistory(username)
```
##### Get A Certain Number Of Recent Balance History
```
recentBalanceHistory, err := api.GetRecentBalanceHistory(uesrname, numOfHistory)
```
##### Get Balance History In The Range Of Index [from, to] Inclusively
```
rangedBalanceHistory, err := api.GetBalanceHistoryFromTo(username, from, to)
```
##### Get Balance History From A Certain Bucket
```
bucketBalanceHistory, err := api.GetBalanceHistory(username, bucketIndex)
```
##### Get Granted Public Key
```
grantPubKey, err := api.GetGrantPubKey(username, pubKeyHex)
```
##### Get Reward
```
reward, err := api.GetReward(username)
```
##### Get Reward At A Certain Block Height
```
reward, err := api.GetRewardAtHeight(username, height)
```
##### Get All Reward History From All Buckets
```
allRewardHistory, err := api.GetAllRewardHistory(username)
```
##### Get A Certain Number Of Recent Reward History
```
recentRewardHistory, err := api.GetRecentRewardHistory(username, numOfHistory) 
```
##### Get Reward History In The Range Of Index [from, to] Inclusively
```
rangedRewardHistory, err := api.GetRewardHistoryFromTo(username, from, to)
```
##### Get Reward History From A Certain Bucket
```
bucketRewardHistory, err := api.GetRewardHistory(username, bucketIndex)
```
##### Get Donation Relationship
```
relationship, err := api.GetRelationship(me, other)
```
##### Get Follower Meta
```
followerMeta, err := api.GetFollowerMeta(me, myFollower)
```
##### Get Following Meta
```
followingMeta, err := api.GetFollowingMeta(me, myFollowing)
```
##### Get All Granted Public Keys
```
pubKeyToGrantPubKeyMap, err := api.GetAllGrantPubKeys(username)
```
##### Get All Donation Relationships 
```
userToRelationshipMap, err := api.GetAllRelationships(username)
```
##### Get All Follower Meta
```
followerToMetaMap, err := api.GetAllFollowerMeta(username)
```
##### Get All Following Meta
```
followingToMetaMap, err := api.GetAllFollowingMeta(username)
```

#### Developer
##### Get Developer 
```
developer, err := api.GetDeveloper(developerName)
```
##### Get All Developers
```
devevlopers, err := api.GetDevelopers()
```

#### Infrastructure
##### Get Infra Provider
```
infraProvider, err := api.GetInfraProvider(providerName)
```
##### Get All Infra Providers
```
infraProviders, err := api.GetInfraProviders()
```

#### Blockchain Parameters
##### Get Evaluate Of Content Value Param
```
p, err := api.GetEvaluateOfContentValueParam()
```
##### Get Global Allocation Param
```
p, err := api.GetGlobalAllocationParam()
```
##### Get Infra Internal Allocation Param
```
p, err := api.GetInfraInternalAllocation()
```
##### Get Developer Param
```
p, err := api.GetDeveloperParam()
```
##### Get Vote Param
```
p, err := api.GetVoteParam()
```
##### Get Proposal Param
```
p, err := api.GetProposalParam()
```
##### Get Validator Param
```
p, err := api.GetValidatorParam()
```
##### Get Coin Day Param
```
p, err := api.GetCoinDayParam()
```
##### Get Bandwidth Param
```
p, err := api.GetBandwidthParam()
```
##### Get Account Param
```
p, err := api.GetAccountParam()
```
##### Get Post Param
```
p, err := api.GetPostParam()
```

#### Post
##### Get PostInfo
```
postInfo, err := api.GetPostInfo(author, postID)
```
##### Get PostMeta
```
postMeta, err := api.GetPostMeta(author, postID)
```
##### Get Post Comment
```
comment, err := api.GetPostComment(author, postID, commentPermlink)
```
##### Get Post View
```
view, err := api.GetPostView(author, postID, viewUser)
```
##### Get Post Donations
```
donations, err := api.GetPostDonations(author, postID, donateUser)
```
##### Get Post ReportOrUpvote
```
reportOrUpvote, err := api.GetPostReportOrUpvote(author, postID, user)
```
##### Get User All Posts
```
permlinkToPostMap, err := api.GetUserAllPosts(username)
```
##### Get Post All Comments
```
permlinkToCommentMap, err := api.GetPostAllComments(author, postID)
```
##### Get Post All Views
```
userToViewMap, err := api.GetPostAllViews(author, postID)
```
##### Get Post All Donations
```
userToDonationsMap, err := api.GetPostAllDonations(author, postID)
```
##### Get Post All ReportOrUpvotes
```
userToReportOrUpvoteMap, err := api.GetPostAllReportOrUpvotes(author, postID)
```

#### Proposal
##### Get Proposal List
```
proposalList, err := api.GetProposalList()
```
##### Get Proposal 
```
proposal, err := api.GetProposal(proposalID)
```
##### Get Ongoing Proposals
```
ongoingProposals, err := api.GetOngoingProposal()
```
##### Get Expired Proposals
```
expiredProposals, err := api.GetExpiredProposal()
```
##### Get Next Proposal ID
```
nextProposalID, err := api.GetNextProposalID()
```

#### Block
##### Get Block
```
block, err := api.GetBlock(height)
```
##### Get Block Status
```
blockStatus, err := api.GetBlockStatus()
```

#### Validator
##### Get Validator
```
validator, err := api.GetValidator()
```
##### Get All Validators
```
validators, err := api.GetAllValidators()
```

### Broadcast
#### Account
##### Register A New User
```
seq, err := api.GetSeqNumber(referrer)
err = api.Register(referrer, registerFee, newUsername, newUserResetPubHex, newUserTxPubHex, newUserAppPubHex, referrerTxPrivKey, seq)
```
##### Transfer LINO Between two users
```
seq, err := api.GetSeqNumber(sender)
err = api.Transfer(sender, receiver, amount, memo, privKeyHex, seq)
```
##### Follow 
```
seq, err := api.GetSeqNumber(follower)
err = api.Follow(follower, followee, privKeyHex, seq)
```
##### Unfollow 
```
seq, err := api.GetSeqNumber(follower)
err = api.Follow(follower, followee, privKeyHex, seq)
```
##### Claim Reward
```
seq, err := api.GetSeqNumber(username)
err = api.Follow(username, privKeyHex, seq)
```
##### Update Account
```
seq, err := api.GetSeqNumber(username)
err = api.UpdateAccount(username, jsonMeta, privKeyHex, seq)
```
##### Recover 
```
seq, err := api.GetSeqNumber(username)
err = api.Recover(username, newResetPubKeyHex, newTransactionPubKeyHex, newAppPubKeyHex, privKeyHex, seq)
```

#### Post
##### Create Post
```
seq, err := api.GetSeqNumber(author)
err = api.CreatePost(author, postID, title, content, parentAuthor, parentPostID, sourceAuthor, sourcePostID, redistributionSplitRate, links, privKeyHex, seq)
```
##### Donate To A Post
```
seq, err := api.GetSeqNumber(username)
err = api.Donate(username, author, amount, postID, fromApp, memo, privKeyHex, seq)
```
##### ReportOrUpvote To A Post
```
seq, err := api.GetSeqNumber(username)
err = api.ReportOrUpvote(username, author, postID, isReport, privKeyHex, seq)
```
##### Delete Post
```
seq, err := api.GetSeqNumber(author)
err = api.DeletePost(author, postID, privKeyHex, seq)
```
##### View A Post
```
seq, err := api.GetSeqNumber(username)
err = api.View(username, author, postID, privKeyHex, seq)
```
##### Update Post
```
seq, err := api.GetSeqNumber(author)
err = api.UpdatePost(author, title, postID, content, redistributionSplitRate, links, privKeyHex, seq)
```

#### Validator
##### Validator Deposit
```
seq, err := api.GetSeqNumber(username)
err = api.ValidatorDeposit(username, deposit, validatorPubKey, link, privKeyHex, seq)
```
##### Validator Withdraw
```
seq, err := api.GetSeqNumber(username)
err = api.ValidatorWithdraw(username, amount, privKeyHex, seq)
```
##### Validator Revoke
```
seq, err := api.GetSeqNumber(username)
err = api.ValidatorRevoke(username, privKeyHex, seq)
```

#### Vote
##### Voter Deposit
```
seq, err := api.GetSeqNumber(username)
err = api.VoterDeposit(username, deposit, privKeyHex, seq)
```
##### Voter Withdraw
```
seq, err := api.GetSeqNumber(username)
err = api.VoterWithdraw(username, amount, privKeyHex, seq)
```
##### Voter Revoke
```
seq, err := api.GetSeqNumber(username)
err = api.VoterRevoke(username, privKeyHex, seq)
```
##### Delegate To Voter
```
seq, err := api.GetSeqNumber(delegator)
err = api.Delegate(delegator, voter, amount, privKeyHex, seq)
```
##### Delegator Withdraw
```
seq, err := api.GetSeqNumber(delegator)
err = api.DelegatorWithdraw(delegator, voter, amount, privKeyHex, seq)
```
##### RevokeDelegation
```
seq, err := api.GetSeqNumber(delegator)
err = api.RevokeDelegation(delegator, voter, privKeyHex, seq)
```

#### Developer
##### Developer Register
```
seq, err := api.GetSeqNumber(username)
err = api.DeveloperRegister(username, deposit, website, description, appMetaData, privKeyHex, seq)
```
##### DeveloperUpdate
```
seq, err := api.GetSeqNumber(username)
err = api.DeveloperUpdate(username, website, description, appMetaData, privKeyHex, seq)
```
##### DeveloperRevoke
```
seq, err := api.GetSeqNumber(username)
err = api.DeveloperRevoke(username, privKeyHex, seq)
```
##### Grant Permission
```
seq, err := api.GetSeqNumber(username)
err = api.GrantPermission(username, authorizedApp, validityPeriodSec, grantLevel, privKeyHex, seq)
```
##### Pre Authorization Permission
```
seq, err := api.GetSeqNumber(username)
err = api.PreAuthorizationPermission(username, authorizedApp, validityPeriodSec, amount, privKeyHex, seq)
```
##### Revoke Permission
```
seq, err := api.GetSeqNumber(username)
err = api.RevokePermission(username, pubKeyHex, privKeyHex, seq)
```

#### Infra
##### Infra Provider Report
```
seq, err := api.GetSeqNumber(username)
err = api.ProviderReport(username, usage, privKeyHex, seq)
```

#### Proposal
##### Change Evaluate Of Content Value Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeEvaluateOfContentValueParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Global Allocation Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeGlovalAllocationParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Infra Internal Allocation Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeInfraInternalAllocationParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Vote Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeVoteParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Proposal Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeProposalParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Developer Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeDeveloperParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Validator Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeValidatorParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Bandwidth Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeBandwidthParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Account Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangeAccountParam(creator, parameter, reason, privKeyHex, seq)
```
##### Change Post Param
```
seq, err := api.GetSeqNumber(creator)
err = api.ChangePostParam(creator, parameter, reason, privKeyHex, seq)
```
##### Delete Post Content
```
seq, err := api.GetSeqNumber(creator)
err = api.DeletePostContent(creator, postAuthor, postID, reason, privKeyHex, seq)
```
##### Vote Proposal
```
seq, err := api.GetSeqNumber(voter)
err = api.VoteProposal(voter, proposalID, result, privKeyHex, seq)
```
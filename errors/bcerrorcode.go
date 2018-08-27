package errors

type BCCodeType uint16

const (
	// See https://github.com/cosmos/cosmos-sdk/issues/766
	LinoErrorCodeSpace = 11

	// Lino common errors reserve 100 ~ 149
	CodeInvalidUsername     BCCodeType = 100
	CodeAccountNotFound     BCCodeType = 101
	CodeFailedToMarshal     BCCodeType = 102
	CodeFailedToUnmarshal   BCCodeType = 103
	CodeIllegalWithdraw     BCCodeType = 104
	CodeInsufficientDeposit BCCodeType = 105
	CodeInvalidCoin         BCCodeType = 106
	CodePostNotFound        BCCodeType = 107
	CodeDeveloperNotFound   BCCodeType = 108
	CodeInvalidCoins        BCCodeType = 109

	// Lino authenticate errors reserve 150 ~ 199
	CodeIncorrectStdTxType   BCCodeType = 150
	CodeNoSignatures         BCCodeType = 151
	CodeUnknownMsgType       BCCodeType = 152
	CodeWrongNumberOfSigners BCCodeType = 153
	CodeInvalidSequence      BCCodeType = 154
	CodeUnverifiedBytes      BCCodeType = 155

	// ABCI Response Codes
	CodeGenesisFailed BCCodeType = 200

	// // Lino register handler errors reserve 300 ~ 309.
	// CodeAccRegisterFailed BCCodeType = 302
	// CodeUsernameNotFound  BCCodeType = 303

	// Lino account errors reserve 300 ~ 399
	CodeRewardNotFound                     BCCodeType = 300
	CodeAccountMetaNotFound                BCCodeType = 301
	CodeAccountInfoNotFound                BCCodeType = 302
	CodeAccountBankNotFound                BCCodeType = 303
	CodePendingStakeQueueNotFound          BCCodeType = 304
	CodeGrantPubKeyNotFound                BCCodeType = 305
	CodeFailedToMarshalAccountInfo         BCCodeType = 306
	CodeFailedToMarshalAccountBank         BCCodeType = 307
	CodeFailedToMarshalAccountMeta         BCCodeType = 308
	CodeFailedToMarshalFollowerMeta        BCCodeType = 309
	CodeFailedToMarshalFollowingMeta       BCCodeType = 310
	CodeFailedToMarshalReward              BCCodeType = 311
	CodeFailedToMarshalPendingStakeQueue   BCCodeType = 312
	CodeFailedToMarshalGrantPubKey         BCCodeType = 313
	CodeFailedToMarshalRelationship        BCCodeType = 314
	CodeFailedToMarshalBalanceHistory      BCCodeType = 315
	CodeFailedToUnmarshalAccountInfo       BCCodeType = 316
	CodeFailedToUnmarshalAccountBank       BCCodeType = 317
	CodeFailedToUnmarshalAccountMeta       BCCodeType = 318
	CodeFailedToUnmarshalReward            BCCodeType = 319
	CodeFailedToUnmarshalPendingStakeQueue BCCodeType = 320
	CodeFailedToUnmarshalGrantPubKey       BCCodeType = 321
	CodeFailedToUnmarshalRelationship      BCCodeType = 322
	CodeFailedToUnmarshalBalanceHistory    BCCodeType = 323
	CodeFolloweeNotFound                   BCCodeType = 324
	CodeFollowerNotFound                   BCCodeType = 325
	CodeReceiverNotFound                   BCCodeType = 326
	CodeSenderNotFound                     BCCodeType = 327
	CodeReferrerNotFound                   BCCodeType = 328
	CodeAddSavingCoinWithFullStake         BCCodeType = 329
	CodeAddSavingCoin                      BCCodeType = 330
	CodeInvalidMemo                        BCCodeType = 331
	CodeInvalidJSONMeta                    BCCodeType = 332
	CodeCheckResetKey                      BCCodeType = 333
	CodeCheckTransactionKey                BCCodeType = 334
	CodeCheckGrantAppKey                   BCCodeType = 335
	CodeCheckAuthenticatePubKeyOwner       BCCodeType = 336
	CodeGrantKeyExpired                    BCCodeType = 337
	CodeGrantKeyNoLeftTimes                BCCodeType = 338
	CodeGrantKeyMismatch                   BCCodeType = 339
	CodeAppGrantKeyMismatch                BCCodeType = 340
	CodeGetResetKey                        BCCodeType = 341
	CodeGetTransactionKey                  BCCodeType = 342
	CodeGetAppKey                          BCCodeType = 343
	CodeGetSavingFromBank                  BCCodeType = 344
	CodeGetSequence                        BCCodeType = 345
	CodeGetLastReportOrUpvoteAt            BCCodeType = 346
	CodeUpdateLastReportOrUpvoteAt         BCCodeType = 347
	CodeGetFrozenMoneyList                 BCCodeType = 348
	CodeIncreaseSequenceByOne              BCCodeType = 349
	CodeGrantTimesExceedsLimitation        BCCodeType = 350
	CodeUnsupportGrantLevel                BCCodeType = 351
	CodeRevokePermissionLevelMismatch      BCCodeType = 352
	CodeCheckUserTPSCapacity               BCCodeType = 353
	CodeAccountTPSCapacityNotEnough        BCCodeType = 354
	CodeAccountSavingCoinNotEnough         BCCodeType = 355
	CodeAccountAlreadyExists               BCCodeType = 356
	CodeRegisterFeeInsufficient            BCCodeType = 357
	CodeFailedToMarshalRewardHistory       BCCodeType = 358
	CodeFailedToUnmarshalRewardHistory     BCCodeType = 359
	CodeGetLastPostAt                      BCCodeType = 360
	CodeUpdateLastPostAt                   BCCodeType = 361

	// Lino post errors reserve 400 ~ 499
	CodePostMetaNotFound                     BCCodeType = 400
	CodePostReportOrUpvoteNotFound           BCCodeType = 401
	CodePostCommentNotFound                  BCCodeType = 402
	CodePostViewNotFound                     BCCodeType = 403
	CodePostDonationNotFound                 BCCodeType = 404
	CodeFailedToMarshalPostInfo              BCCodeType = 405
	CodeFailedToMarshalPostMeta              BCCodeType = 406
	CodeFailedToMarshalPostReportOrUpvote    BCCodeType = 407
	CodeFailedToMarshalPostComment           BCCodeType = 408
	CodeFailedToMarshalPostView              BCCodeType = 409
	CodeFailedToMarshalPostDonations         BCCodeType = 410
	CodeFailedToUnmarshalPostInfo            BCCodeType = 411
	CodeFailedToUnmarshalPostMeta            BCCodeType = 412
	CodeFailedToUnmarshalPostReportOrUpvote  BCCodeType = 413
	CodeFailedToUnmarshalPostComment         BCCodeType = 414
	CodeFailedToUnmarshalPostView            BCCodeType = 415
	CodeFailedToUnmarshalPostDonations       BCCodeType = 416
	CodePostAlreadyExist                     BCCodeType = 417
	CodeInvalidPostRedistributionSplitRate   BCCodeType = 418
	CodeDonatePostIsDeleted                  BCCodeType = 419
	CodeCannotDonateToSelf                   BCCodeType = 420
	CodeProcessSourceDonation                BCCodeType = 421
	CodeProcessDonation                      BCCodeType = 422
	CodeUpdatePostIsDeleted                  BCCodeType = 423
	CodeReportOrUpvoteTooOften               BCCodeType = 424
	CodeReportOrUpvoteAlreadyExist           BCCodeType = 425
	CodeNoPostID                             BCCodeType = 426
	CodePostIDTooLong                        BCCodeType = 427
	CodeNoAuthor                             BCCodeType = 428
	CodeNoUsername                           BCCodeType = 429
	CodeCommentAndRepostConflict             BCCodeType = 430
	CodePostTitleExceedMaxLength             BCCodeType = 431
	CodePostContentExceedMaxLength           BCCodeType = 432
	CodeRedistributionSplitRateLengthTooLong BCCodeType = 433
	CodeIdentifierLengthTooLong              BCCodeType = 434
	CodeURLLengthTooLong                     BCCodeType = 435
	CodeTooManyURL                           BCCodeType = 436
	CodeInvalidTarget                        BCCodeType = 437
	CodeCreatePostSourceInvalid              BCCodeType = 438
	CodeGetSourcePost                        BCCodeType = 439
	CodePostTooOften                         BCCodeType = 440

	// Lino validator errors reserve 500 ~ 599
	CodeValidatorNotFound              BCCodeType = 500
	CodeValidatorListNotFound          BCCodeType = 501
	CodeFailedToMarshalValidator       BCCodeType = 502
	CodeFailedToMarshalValidatorList   BCCodeType = 503
	CodeFailedToUnmarshalValidator     BCCodeType = 504
	CodeFailedToUnmarshalValidatorList BCCodeType = 505
	CodeUnbalancedAccount              BCCodeType = 506
	CodeValidatorPubKeyAlreadyExist    BCCodeType = 507

	// Lino global errors reserve 600 ~ 699
	CodeInfraInflationCoinConversion     BCCodeType = 600
	CodeContentCreatorCoinConversion     BCCodeType = 601
	CodeDeveloperCoinConversion          BCCodeType = 602
	CodeValidatorCoinConversion          BCCodeType = 603
	CodeGlobalMetaNotFound               BCCodeType = 604
	CodeInflationPoolNotFound            BCCodeType = 605
	CodeGlobalConsumptionMetaNotFound    BCCodeType = 606
	CodeGlobalTPSNotFound                BCCodeType = 607
	CodeFailedToMarshalTimeEventList     BCCodeType = 608
	CodeFailedToMarshalGlobalMeta        BCCodeType = 609
	CodeFailedToMarshalInflationPoll     BCCodeType = 610
	CodeFailedToMarshalConsumptionMeta   BCCodeType = 611
	CodeFailedToMarshalTPS               BCCodeType = 612
	CodeFailedToUnmarshalTimeEventList   BCCodeType = 613
	CodeFailedToUnmarshalGlobalMeta      BCCodeType = 614
	CodeFailedToUnmarshalInflationPool   BCCodeType = 615
	CodeFailedToUnmarshalConsumptionMeta BCCodeType = 616
	CodeFailedToUnmarshalTPS             BCCodeType = 617
	CodeRegisterExpiredEvent             BCCodeType = 618

	// Vote errors reserve 700 ~ 799
	CodeVoterNotFound                  BCCodeType = 700
	CodeVoteNotFound                   BCCodeType = 701
	CodeReferenceListNotFound          BCCodeType = 702
	CodeDelegationNotFound             BCCodeType = 703
	CodeFailedToMarshalVoter           BCCodeType = 704
	CodeFailedToMarshalVote            BCCodeType = 705
	CodeFailedToMarshalDelegation      BCCodeType = 706
	CodeFailedToMarshalReferenceList   BCCodeType = 707
	CodeFailedToUnmarshalVoter         BCCodeType = 708
	CodeFailedToUnmarshalVote          BCCodeType = 709
	CodeFailedToUnmarshalDelegation    BCCodeType = 710
	CodeFailedToUnmarshalReferenceList BCCodeType = 711
	CodeValidatorCannotRevoke          BCCodeType = 712
	CodeVoteAlreadyExist               BCCodeType = 713

	// Lino infra errors reserve 800 ~ 899
	CodeInfraProviderNotFound              BCCodeType = 800
	CodeInfraProviderListNotFound          BCCodeType = 801
	CodeFailedToMarshalInfraProvider       BCCodeType = 802
	CodeFailedToMarshalInfraProviderList   BCCodeType = 803
	CodeFailedToUnmarshalInfraProvider     BCCodeType = 804
	CodeFailedToUnmarshalInfraProviderList BCCodeType = 805
	CodeInvalidUsage                       BCCodeType = 806

	// Lino developer errors reserve 900 ~ 999
	CodeDeveloperListNotFound          BCCodeType = 900
	CodeFailedToMarshalDeveloper       BCCodeType = 901
	CodeFailedToMarshalDeveloperList   BCCodeType = 902
	CodeFailedToUnmarshalDeveloper     BCCodeType = 903
	CodeFailedToUnmarshalDeveloperList BCCodeType = 904
	CodeDeveloperAlreadyExist          BCCodeType = 905
	CodeInsufficientDeveloperDeposit   BCCodeType = 906
	CodeInvalidAuthorizedApp           BCCodeType = 907
	CodeInvalidValidityPeriod          BCCodeType = 908
	CodeGrantPermissionTooHigh         BCCodeType = 909
	CodeInvalidWebsite                 BCCodeType = 910
	CodeInvalidDescription             BCCodeType = 911
	CodeInvalidAppMetadata             BCCodeType = 912

	// Param errors reserve 1000 ~ 1099
	CodeParamHolderGenesisError                       BCCodeType = 1000
	CodeDeveloperParamNotFound                        BCCodeType = 1001
	CodeValidatorParamNotFound                        BCCodeType = 1002
	CodeCoinDayParamNotFound                          BCCodeType = 1003
	CodeBandwidthParamNotFound                        BCCodeType = 1004
	CodeAccountParamNotFound                          BCCodeType = 1005
	CodeVoteParamNotFound                             BCCodeType = 1006
	CodeProposalParamNotFound                         BCCodeType = 1007
	CodeGlobalAllocationParamNotFound                 BCCodeType = 1008
	CodeInfraAllocationParamNotFound                  BCCodeType = 1009
	CodePostParamNotFound                             BCCodeType = 1010
	CodeInvalidaParameter                             BCCodeType = 1011
	CodeEvaluateOfContentValueParamNotFound           BCCodeType = 1012
	CodeFailedToUnmarshalGlobalAllocationParam        BCCodeType = 1013
	CodeFailedToUnmarshalPostParam                    BCCodeType = 1014
	CodeFailedToUnmarshalValidatorParam               BCCodeType = 1015
	CodeFailedToUnmarshalEvaluateOfContentValueParam  BCCodeType = 1016
	CodeFailedToUnmarshalInfraInternalAllocationParam BCCodeType = 1017
	CodeFailedToUnmarshalDeveloperParam               BCCodeType = 1018
	CodeFailedToUnmarshalVoteParam                    BCCodeType = 1019
	CodeFailedToUnmarshalProposalParam                BCCodeType = 1020
	CodeFailedToUnmarshalCoinDayParam                 BCCodeType = 1021
	CodeFailedToUnmarshalBandwidthParam               BCCodeType = 1022
	CodeFailedToUnmarshalAccountParam                 BCCodeType = 1023
	CodeFailedToMarshalGlobalAllocationParam          BCCodeType = 1024
	CodeFailedToMarshalPostParam                      BCCodeType = 1025
	CodeFailedToMarshalValidatorParam                 BCCodeType = 1026
	CodeFailedToMarshalEvaluateOfContentValueParam    BCCodeType = 1027
	CodeFailedToMarshalInfraInternalAllocationParam   BCCodeType = 1028
	CodeFailedToMarshalDeveloperParam                 BCCodeType = 1029
	CodeFailedToMarshalVoteParam                      BCCodeType = 1030
	CodeFailedToMarshalProposalParam                  BCCodeType = 1031
	CodeFailedToMarshalCoinDayParam                   BCCodeType = 1032
	CodeFailedToMarshalBandwidthParam                 BCCodeType = 1033
	CodeFailedToMarshalAccountParam                   BCCodeType = 1034

	// Proposal errors reserve 1100 ~ 1199
	CodeOngoingProposalNotFound         BCCodeType = 1100
	CodeCensorshipPostNotFound          BCCodeType = 1101
	CodeProposalNotFound                BCCodeType = 1102
	CodeProposalListNotFound            BCCodeType = 1103
	CodeNextProposalIDNotFound          BCCodeType = 1104
	CodeFailedToMarshalProposal         BCCodeType = 1105
	CodeFailedToMarshalProposalList     BCCodeType = 1106
	CodeFailedToMarshalNextProposalID   BCCodeType = 1107
	CodeFailedToUnmarshalProposal       BCCodeType = 1108
	CodeFailedToUnmarshalProposalList   BCCodeType = 1109
	CodeFailedToUnmarshalNextProposalID BCCodeType = 1110
	CodeCensorshipPostIsDeleted         BCCodeType = 1111
	CodeNotOngoingProposal              BCCodeType = 1112
	CodeIncorrectProposalType           BCCodeType = 1113
	CodeInvalidPermlink                 BCCodeType = 1114
	CodeInvalidLink                     BCCodeType = 1115
	CodeIllegalParameter                BCCodeType = 1116
	CodeReasonTooLong                   BCCodeType = 1117
)

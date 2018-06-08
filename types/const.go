package types

const (
	BroadcastMaxTries = 10
	InvalidSeqErrCode = 3

	// as defined by a julian year of 365.25 days
	HoursPerYear = 8766

	// as defined by a julian year of 365.25 days
	MinutesPerYear = HoursPerYear * 60

	// as defined by a julian year of 365.25 days
	MinutesPerMonth = MinutesPerYear / 12

	BalanceHistoryIntervalTime = MinutesPerMonth * 60
)

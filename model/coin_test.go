package model

import (
	"math"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	linotypes "github.com/lino-network/lino/types"
)

var (
	// LowerBoundRat - the lower bound of Rat
	LowerBoundRat = NewDecFromRat(1, Decimals)
	// UpperBoundRat - the upper bound of Rat
	UpperBoundRat = sdk.NewDec(math.MaxInt64 / Decimals)
)

const (
	Decimals = 100000
)

func TestCoinToLNO(t *testing.T) {
	testCases := map[string]struct {
		inputLino     string
		expectCoinStr string
		expectLino    string
	}{
		"lino without 0, coin with 5 zeros": {
			inputLino:     "123",
			expectCoinStr: "12300000",
			expectLino:    "123",
		},
		"lino with some 0, coin with some zeros": {
			inputLino:     "100.00023",
			expectCoinStr: "10000023",
			expectLino:    "100.00023",
		},
		"lino with one 0, coin with more than 5 zeros": {
			inputLino:     "1230",
			expectCoinStr: "123000000",
			expectLino:    "1230",
		},
		"lino with one digit, coin with less than 5 zeros": {
			inputLino:     "12.3",
			expectCoinStr: "1230000",
			expectLino:    "12.3",
		},
		"lino with three digits, coin with less than 5 zeros": {
			inputLino:     "0.123",
			expectCoinStr: "12300",
			expectLino:    "0.123",
		},
		"lino with five digits, coin with no zero": {
			inputLino:     "0.00123",
			expectCoinStr: "123",
			expectLino:    "0.00123",
		},
		"lino with 3 zero, coin with no zero": {
			inputLino:     "100082.92819",
			expectCoinStr: "10008292819",
			expectLino:    "100082.92819",
		},
	}

	for testName, tc := range testCases {
		coin, err := linotypes.LinoToCoin(tc.inputLino)
		if err != nil {
			t.Errorf("%s: failed to convert lino to coin, got err %v", testName, err)
		}

		if coin.Amount.String() != tc.expectCoinStr {
			t.Errorf("%s: diff coin amount, got %v, want %v", testName, coin.Amount.String(), tc.expectCoinStr)
		}

		got := CoinToLNO(coin)
		if got != tc.expectLino {
			t.Errorf("%s: diff lino, got %v, want %v", testName, got, tc.expectLino)
		}
	}
}

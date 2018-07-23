package model

import (
	"errors"
	"math"
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	LowerBoundRat = big.NewRat(1, Decimals)
	UpperBoundRat = big.NewRat(math.MaxInt64/Decimals, 1)
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
	}

	for testName, tc := range testCases {
		coin, err := LinoToCoin(tc.inputLino)
		if err != nil {
			t.Errorf("%s: failed to convert lino to coin, got err %v", testName, err)
		}

		if coin.Amount.String() != tc.expectCoinStr {
			t.Errorf("%s: diff coin amount, got %v, want %v", testName, coin.Amount.String(), tc.expectCoinStr)
		}

		got := coin.CoinToLNO()
		if got != tc.expectLino {
			t.Errorf("%s: diff lino, got %v, want %v", testName, got, tc.expectLino)
		}
	}
}

//
// helper function
//

func NewCoinFromInt64(amount int64) Coin {
	return Coin{sdk.NewInt(amount)}
}

func NewCoinFromBigInt(amount *big.Int) Coin {
	sdkInt := sdk.NewIntFromBigInt(amount)
	return Coin{sdkInt}
}

func LinoToCoin(lino string) (Coin, error) {
	num, success := new(big.Rat).SetString(lino)
	if !success {
		return NewCoinFromInt64(0), errors.New("illegal")
	}
	if num.Cmp(UpperBoundRat) > 0 {
		return NewCoinFromInt64(0), errors.New("overflow")
	}
	if num.Cmp(LowerBoundRat) < 0 {
		return NewCoinFromInt64(0), errors.New("underflow")
	}
	return RatToCoin(sdk.Rat{new(big.Rat).Mul(num, big.NewRat(Decimals, 1))}), nil
}

func RatToCoin(rat sdk.Rat) Coin {
	return NewCoinFromBigInt(rat.EvaluateBig())
}

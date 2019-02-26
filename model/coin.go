package model

import (
	"fmt"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Coin is the same struct used in Lino blockchain.
type Coin struct {
	Amount sdk.Int `json:"amount"`
}

// NewCoinFromString - return string amount of Coin
func NewCoinFromString(amount string) (Coin, bool) {
	res, ok := sdk.NewIntFromString(amount)
	return Coin{res}, ok
}

// NewCoinFromBigInt - return big.Int amount of Coin
func NewCoinFromBigInt(amount *big.Int) Coin {
	sdkInt := sdk.NewIntFromBigInt(amount)
	return Coin{sdkInt}
}

// DecToCoin - convert sdk.Dec to LNO coin
// XXX(yumin): the unit of @p rat must be coin.
func DecToCoin(rat sdk.Dec) Coin {
	return NewCoinFromBigInt(rat.RoundInt().BigInt())
}

func (c Coin) CoinToLNO() string {
	amountStr := "00000" + c.Amount.String()
	length := len(amountStr)
	converted := fmt.Sprintf("%s.%s", amountStr[:length-5], amountStr[length-5:])
	converted = strings.Trim(converted, "0")
	if converted[0] == '.' {
		converted = "0" + converted
	}
	if converted[len(converted)-1] == '.' {
		converted = converted[:len(converted)-1]
	}
	return converted
}

// IsZero returns if this represents no money
func (coin Coin) IsZero() bool {
	return coin.Amount.Sign() == 0
}

// IsGT returns true if the receiver is greater value
func (coin Coin) IsGT(other Coin) bool {
	return coin.Amount.GT(other.Amount)
}

// IsGTE returns true if they are the same type and the receiver is
// an equal or greater value
func (coin Coin) IsGTE(other Coin) bool {
	return coin.Amount.GT(other.Amount) || coin.Amount.Equal(other.Amount)
}

// IsEqual returns true if the two sets of Coins have the same value
func (coin Coin) IsEqual(other Coin) bool {
	return coin.Amount.Equal(other.Amount)
}

// IsPositive returns true if coin amount is positive
func (coin Coin) IsPositive() bool {
	return coin.Amount.Sign() > 0
}

// IsNotNegative returns true if coin amount is not negative
func (coin Coin) IsNotNegative() bool {
	return coin.Amount.Sign() >= 0
}

// Adds amounts of two coins with same denom
func (coin Coin) Plus(coinB Coin) Coin {
	r := coin.Amount.Add(coinB.Amount)
	return Coin{r}
}

// Subtracts amounts of two coins with same denom
func (coin Coin) Minus(coinB Coin) Coin {
	r := coin.Amount.Sub(coinB.Amount)
	return Coin{r}
}

// SDKCoin is the same struct used in cosmos-sdk.
type SDKCoin struct {
	Denom  string `json:"denom"`
	Amount int64  `json:"amount"`
}

type SDKCoins []SDKCoin

type Rat struct {
	*big.Rat `json:"rat"`
}

// MarshalAmino wraps r.MarshalText().
func (r Rat) MarshalAmino() (string, error) {
	if r.Rat == nil {
		r.Rat = new(big.Rat)
	}
	bz, err := r.Rat.MarshalText()
	return string(bz), err
}

// UnmarshalAmino requires a valid JSON string - strings quotes and calls UnmarshalText
func (r *Rat) UnmarshalAmino(text string) (err error) {
	tempRat := big.NewRat(0, 1)
	err = tempRat.UnmarshalText([]byte(text))
	if err != nil {
		return err
	}
	r.Rat = tempRat
	return nil
}

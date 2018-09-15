package model

import (
	"math/big"
	"strings"
)

// Coin is the same struct used in Lino blockchain.
type Coin struct {
	Amount Int `json:"amount"`
}

func NewCoinFromString(amount string) (Coin, bool) {
	res, ok := NewIntFromString(amount)
	return Coin{res}, ok
}

func (c Coin) CoinToLNO() string {
	amountStr := c.Amount.String()

	numZero := strings.Count(amountStr, "0")
	if numZero >= 5 {
		index := len(amountStr) - 5
		return amountStr[:index]
	} else {
		numOfRemainZero := 5 - numZero
		amountStr = strings.TrimRight(amountStr, "0")

		if len(amountStr) > numOfRemainZero {
			index := len(amountStr) - numOfRemainZero
			return amountStr[:index] + "." + amountStr[index:]
		} else {
			numOfAddingZero := numOfRemainZero - len(amountStr)
			res := "0."
			for numOfAddingZero > 0 {
				res += "0"
				numOfAddingZero--
			}

			return res + amountStr
		}
	}

	return ""
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

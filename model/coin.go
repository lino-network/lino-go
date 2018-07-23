package model

import (
	"encoding/json"
	"math/big"
	"strings"
)

// Coin is the same struct used in Lino blockchain.
type Coin struct {
	Amount Int `json:"amount"`
}

type Int struct {
	i *big.Int
}

func (i Int) String() string {
	return i.i.String()
}

// MarshalAmino for custom encoding scheme
func marshalAmino(i *big.Int) (string, error) {
	bz, err := i.MarshalText()
	return string(bz), err
}

// UnmarshalAmino for custom decoding scheme
func unmarshalAmino(i *big.Int, text string) (err error) {
	return i.UnmarshalText([]byte(text))
}

// MarshalJSON for custom encoding scheme
// Must be encoded as a string for JSON precision
func marshalJSON(i *big.Int) ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

// UnmarshalJSON for custom decoding scheme
// Must be encoded as a string for JSON precision
func unmarshalJSON(i *big.Int, bz []byte) error {
	var text string
	err := json.Unmarshal(bz, &text)
	if err != nil {
		return err
	}
	return i.UnmarshalText([]byte(text))
}

// MarshalAmino defines custom encoding scheme
func (i Int) MarshalAmino() (string, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalAmino(i.i)
}

// UnmarshalAmino defines custom decoding scheme
func (i *Int) UnmarshalAmino(text string) error {
	if i.i == nil { // Necessary since default Int initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalAmino(i.i, text)
}

// MarshalJSON defines custom encoding scheme
func (i Int) MarshalJSON() ([]byte, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalJSON(i.i)
}

// UnmarshalJSON defines custom decoding scheme
func (i *Int) UnmarshalJSON(bz []byte) error {
	if i.i == nil { // Necessary since default Int initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalJSON(i.i, bz)
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

// SDKCoin is the same struct used in cosmos-sdk.
type SDKCoin struct {
	Denom  string `json:"denom"`
	Amount int64  `json:"amount"`
}

type SDKCoins []SDKCoin

type Rat struct {
	big.Rat `json:"rat"`
}

// MarshalAmino wraps r.MarshalText().
func (r Rat) MarshalAmino() (string, error) {
	bz, err := (&(r.Rat)).MarshalText()
	return string(bz), err
}

// UnmarshalAmino requires a valid JSON string - strings quotes and calls UnmarshalText
func (r *Rat) UnmarshalAmino(text string) (err error) {
	tempRat := big.NewRat(0, 1)
	err = tempRat.UnmarshalText([]byte(text))
	if err != nil {
		return err
	}
	r.Rat = *tempRat
	return nil
}

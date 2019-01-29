package model

import (
	"encoding/json"
	"math/big"
)

type Int struct {
	I *big.Int
}

// BigInt converts Int to big.Int
func (i Int) BigInt() *big.Int {
	return new(big.Int).Set(i.I)
}

// NewIntFromString constructs Int from string
func NewIntFromString(s string) (res Int, ok bool) {
	i, ok := newIntegerFromString(s)
	if !ok {
		return
	}
	// Check overflow
	if i.BitLen() > 255 {
		ok = false
		return
	}
	return Int{i}, true
}

// IsZero returns true if Int is zero
func (i Int) IsZero() bool {
	return i.I.Sign() == 0
}

// Sign returns sign of Int
func (i Int) Sign() int {
	return i.I.Sign()
}

// Equal compares two Ints
func (i Int) Equal(i2 Int) bool {
	return equal(i.I, i2.I)
}

// GT returns true if first Int is greater than second
func (i Int) GT(i2 Int) bool {
	return gt(i.I, i2.I)
}

// LT returns true if first Int is lesser than second
func (i Int) LT(i2 Int) bool {
	return lt(i.I, i2.I)
}

// Add adds Int from another
func (i Int) Add(i2 Int) (res Int) {
	res = Int{add(i.I, i2.I)}
	// Check overflow
	if res.I.BitLen() > 255 {
		panic("Int overflow")
	}
	return
}

// Sub subtracts Int from another
func (i Int) Sub(i2 Int) (res Int) {
	res = Int{sub(i.I, i2.I)}
	// Check overflow
	if res.I.BitLen() > 255 {
		panic("Int overflow")
	}
	return
}

// Mul multiples two Ints
func (i Int) Mul(i2 Int) (res Int) {
	// Check overflow
	if i.I.BitLen()+i2.I.BitLen()-1 > 255 {
		panic("Int overflow")
	}
	res = Int{mul(i.I, i2.I)}
	// Check overflow if sign of both are same
	if res.I.BitLen() > 255 {
		panic("Int overflow")
	}
	return
}

// Div divides Int with Int
func (i Int) Div(i2 Int) (res Int) {
	// Check division-by-zero
	if i2.I.Sign() == 0 {
		panic("Division by zero")
	}
	return Int{div(i.I, i2.I)}
}

// Neg negates Int
func (i Int) Neg() (res Int) {
	return Int{neg(i.I)}
}

// Return the minimum of the ints
func MinInt(i1, i2 Int) Int {
	return Int{min(i1.BigInt(), i2.BigInt())}
}

func (i Int) String() string {
	return i.I.String()
}

func newIntegerFromString(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 0)
}

func equal(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 0 }

func gt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 1 }

func lt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == -1 }

func add(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

func sub(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Sub(i, i2) }

func mul(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mul(i, i2) }

func div(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Div(i, i2) }

// func mod(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mod(i, i2) }

func neg(i *big.Int) *big.Int { return new(big.Int).Neg(i) }

func min(i *big.Int, i2 *big.Int) *big.Int {
	if i.Cmp(i2) == 1 {
		return new(big.Int).Set(i2)
	}
	return new(big.Int).Set(i)
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
	if i.I == nil { // Necessary since default Uint initialization has i.I as nil
		i.I = new(big.Int)
	}
	return marshalAmino(i.I)
}

// UnmarshalAmino defines custom decoding scheme
func (i *Int) UnmarshalAmino(text string) error {
	if i.I == nil { // Necessary since default Int initialization has i.I as nil
		i.I = new(big.Int)
	}
	return unmarshalAmino(i.I, text)
}

// MarshalJSON defines custom encoding scheme
func (i Int) MarshalJSON() ([]byte, error) {
	if i.I == nil { // Necessary since default Uint initialization has i.I as nil
		i.I = new(big.Int)
	}
	return marshalJSON(i.I)
}

// UnmarshalJSON defines custom decoding scheme
func (i *Int) UnmarshalJSON(bz []byte) error {
	if i.I == nil { // Necessary since default Int initialization has i.I as nil
		i.I = new(big.Int)
	}
	return unmarshalJSON(i.I, bz)
}

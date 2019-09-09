package model

import (
	"fmt"
	"strings"

	linotypes "github.com/lino-network/lino/types"
)

func CoinToLNO(c linotypes.Coin) string {
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

package util

// nolint
import (
	"fmt"
	"regexp"
	"strings"

	linotypes "github.com/lino-network/lino/types"
)

const (
	usernameReCheck        = "^[a-z]([a-z0-9-\\.]){1,19}[a-z0-9]$"
	illegalUsernameReCheck = "^[a-z0-9\\.-]*([-\\.]){2,}[a-z0-9\\.-]*$"
)

func CheckUsername(username string) bool {
	match, err := regexp.MatchString(usernameReCheck, username)
	if err != nil || !match {
		return false
	}

	match, err = regexp.MatchString(illegalUsernameReCheck, username)
	if err != nil || match {
		return false
	}
	return true
}

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

func Min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func GetSignerList(signer string) []linotypes.AccOrAddr {
	return []linotypes.AccOrAddr{linotypes.NewAccOrAddrFromAcc(linotypes.AccountKey(signer))}
}

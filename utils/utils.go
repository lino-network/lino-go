package util

// nolint
import (
	"regexp"
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

func Min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

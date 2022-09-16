package env

import (
	"os"
	"strconv"
)

func IgnoreTokenExpiration() bool {
	ignoreTokenExpiration, err := strconv.ParseBool(os.Getenv("IGNORE_TOKEN_EXPIRATION"))
	if err != nil {
		return false
	}
	return ignoreTokenExpiration
}

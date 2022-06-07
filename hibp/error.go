package hibp

import "strings"

func ignore404Error(err error) bool {
	return strings.Contains(err.Error(), "404 Not Found")
}

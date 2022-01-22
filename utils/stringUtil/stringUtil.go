package stringUtil

import (
	"fmt"
	"strconv"
)

func ParseUint(num string) int {
	parsed, err := strconv.ParseUint(num, 10, 32)
	if nil != err {
		fmt.Println("Parsing uint error: ", err)
		return 0
	}
	return int(parsed)
}

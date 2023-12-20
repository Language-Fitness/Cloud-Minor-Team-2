package helper

import "strconv"

func StringToInt(s string) int {
	intValue, _ := strconv.Atoi(s)
	return intValue
}

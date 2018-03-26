package util

import "strconv"

// StringToInteger .
func StringToInteger(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}

package problem0009

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)

	i, j := 0, len(s)-1
	for ; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

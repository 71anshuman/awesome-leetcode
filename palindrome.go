package main

import (
	"strconv"
)

func main() {
	res := isPalindrome(121)
	println(res)

	res = isPalindrome(1)
	println(res)

	res = isPalindrome(12321)
	println(res)

	res = isPalindrome(123)
	println(res)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)

	i, j := 0, len(s)-1
	for ; i <= j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

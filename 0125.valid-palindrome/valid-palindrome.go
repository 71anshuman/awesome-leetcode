package problem0125

import (
	"log"
	"regexp"
	"strings"
)

// Slow
func isPalindrome(s string) bool {
	cleanS := clean(s)
	return cleanS == reverse(cleanS)
}

func clean(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	cleanS := reg.ReplaceAllString(s, "")
	return strings.ToLower(cleanS)
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// Faster solution

func isPalindromeFast(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(ch byte) bool {
	if ('a' <= ch && 'z' >= ch) || (ch >= '0' && ch <= '9') {
		return true
	}

	return false
}

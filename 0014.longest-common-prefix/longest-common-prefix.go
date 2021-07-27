package problem0014

func longestCommonPrefix(strs []string) string {
	prefix := ""
	min := getMinLengthStringSlice(strs)
	for i := 0; i < min; i++ {
		c := strs[0][i]

		for _, s := range strs {
			if c != s[i] {
				return prefix
			}
		}
		prefix = prefix + string(c)
	}

	return prefix
}

func getMinLengthStringSlice(stings []string) int {
	min := len(stings[0])
	for _, str := range stings {
		length := len(str)
		if length < min {
			min = length
		}
	}
	if min == 0 {
		return -1
	}

	return min
}

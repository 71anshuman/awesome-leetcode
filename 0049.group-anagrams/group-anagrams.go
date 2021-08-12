package problem0049

func groupAnagrams(strs []string) [][]string {
	hash := make(map[int][]string)
	var res [][]string
	for _, str := range strs {
		encoded := encodeStr(str)

		hash[encoded] = append(hash[encoded], str)
	}

	for _, slice := range hash {
		res = append(res, slice)
	}

	return res
}

var prime = []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}

func encodeStr(s string) int {
	res := 1
	for i := range s {
		res *= prime[s[i]-'a']
	}
	return res
}

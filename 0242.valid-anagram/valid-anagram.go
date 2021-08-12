package problem0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	st := []rune(t)

	hold := make(map[rune]int, len(sr))

	for i := range sr {
		hold[sr[i]]++
		hold[st[i]]--
	}

	for _, n := range hold {
		if n != 0 {
			return false
		}
	}

	return true
}

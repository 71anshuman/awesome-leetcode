package problem0001

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	var res []int

	for i, val := range nums {
		search := target - val
		v, ok := seen[search]

		if ok {
			res = append(res, i, v)
			break
		}

		_, present := seen[val]

		if !present {
			seen[val] = i
		}
	}

	return res
}

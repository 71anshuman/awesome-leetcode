package problem0011

func maxArea(height []int) int {

	max := 0
	i, j := 0, len(height)-1
	for i < j {
		hi, hj := height[i], height[j]
		h := min(hi, hj)

		a := h * (j - i)
		if a > max {
			max = a
		}

		if hi < hj {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(hi, hj int) int {
	if hi <= hj {
		return hi
	}
	return hj
}

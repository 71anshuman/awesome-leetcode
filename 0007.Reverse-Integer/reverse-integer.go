package problem0007

import "math"

func reverse(x int) int {
	sign := 1
	sum := 0

	//Handling negative numbers
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	for x > 0 {
		// Get the end of x
		rem := x % 10
		// Put x at the beginning of sum
		sum = sum*10 + rem
		// x remove the end
		x = x / 10
	}

	// restore the symbol of x to sum
	sum = sign * sum

	// Deal with the overflow problem of sum
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}

	return sum
}

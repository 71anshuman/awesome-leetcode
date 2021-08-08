package problem0234

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{
	{[]int{1, 2, 3, 2, 1}, true},
	{[]int{1, 3, 2, 1}, false},
}

func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("%v\n", tc)
		head := ints2List(tc.nums)
		ast.Equal(tc.ans, isPalindrome(head), "%v", tc)
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := ints2List(tc.nums)
			isPalindrome(head)
		}
	}
}

func ints2List(nums []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

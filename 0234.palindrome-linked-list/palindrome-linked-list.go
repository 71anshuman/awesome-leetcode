package problem0234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0, 64)

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	i, j := 0, len(nums)-1

	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}

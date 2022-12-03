package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) (result int) {
	set := make([]bool, 1e4+1)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}

	for ; head != nil; head = head.Next {
		if !set[head.Val] {
			continue
		}

		result++
		for ; head != nil && set[head.Val]; head = head.Next {
		}

		if head == nil {
			break
		}
	}

	return result
}

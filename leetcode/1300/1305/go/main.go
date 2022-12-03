package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	first, second := &ListNode{}, &ListNode{}
	_, firstLength := toList(root1, first)
	_, secondLength := toList(root2, second)

	result := make([]int, 0, firstLength+secondLength)
	first, second = first.Next, second.Next
	for first != nil || second != nil {
		if first == nil {
			result = append(result, second.Val)
			second = second.Next
			continue
		}

		if second == nil || first.Val < second.Val {
			result = append(result, first.Val)
			first = first.Next
			continue
		}

		result = append(result, second.Val)
		second = second.Next
	}

	return result
}

func toList(root *TreeNode, head *ListNode) (*ListNode, int) {
	if root == nil {
		return head, 0
	}

	tail, leftLength := toList(root.Left, head)

	tail.Next = &ListNode{Val: root.Val}
	tail = tail.Next

	tail, rightLength := toList(root.Right, tail)

	return tail, leftLength + rightLength + 1
}

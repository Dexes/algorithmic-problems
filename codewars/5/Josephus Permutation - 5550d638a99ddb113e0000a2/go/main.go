package kata

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func Josephus(items []interface{}, k int) []interface{} {
	if k--; k == 0 || len(items) < 2 {
		return items
	}

	var skip int
	result := make([]interface{}, 0, len(items))
	list, listLength := toList(items), len(items)

	for listLength > 0 {
		if skip = k; skip > listLength {
			skip %= listLength
		}

		for i := 0; i < skip; i++ {
			list = list.Next
		}

		result = append(result, list.Next.Val)
		list.Next = list.Next.Next
		listLength--
	}

	return result
}

func toList(items []interface{}) *ListNode {
	head := &ListNode{Val: items[0]}
	tail := head

	for i := 1; i < len(items); i++ {
		tail.Next = &ListNode{Val: items[i]}
		tail = tail.Next
	}

	tail.Next = head
	return tail
}

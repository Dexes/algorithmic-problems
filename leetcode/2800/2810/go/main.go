package main

type node struct {
	val  byte
	next *node
}

func finalString(s string) string {
	head := &node{val: s[0]}
	tail := head
	length := 1
	forward := true

	for i := 1; i < len(s); i++ {
		if s[i] == 'i' {
			forward = !forward
			continue
		}

		length++

		if forward {
			tail.next = &node{val: s[i]}
			tail = tail.next
		} else {
			head = &node{val: s[i], next: head}
		}
	}

	return toString(head, length, forward)
}

func toString(list *node, length int, forward bool) string {
	i, j, step := 0, length, 1
	if !forward {
		i, j, step = length-1, -1, -1
	}

	result := make([]byte, length)
	for ; i != j; i += step {
		result[i] = list.val
		list = list.next
	}

	return string(result)
}

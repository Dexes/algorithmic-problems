package main

import "strings"

type LinkedList struct {
	Head *LinkedNode
	Tail *LinkedNode
	Len  int
}

type LinkedNode struct {
	Value        string
	ContentIndex int
	Next         *LinkedNode
}

func (first *LinkedNode) Compare(second *LinkedNode) int {
	result := strings.Compare(first.Value[first.ContentIndex:], second.Value[second.ContentIndex:])
	if result != 0 {
		return result
	}

	return strings.Compare(first.Value, second.Value)
}

func (list *LinkedList) InsertAndSort(value string, contentIndex int) {
	node := &LinkedNode{Value: value, ContentIndex: contentIndex}
	list.Len++

	if list.Head == nil {
		list.Head, list.Tail = node, node
		return
	}

	if list.Head.Compare(node) == 1 {
		node.Next, list.Head = list.Head, node
		return
	}

	current, next := list.Head, list.Head.Next
	for next != nil {
		if next.Compare(node) == 1 {
			current.Next, node.Next = node, next
			return
		}

		current, next = next, next.Next
	}

	list.Tail, list.Tail.Next = node, node
}

func (list *LinkedList) Insert(value string) {
	node := &LinkedNode{Value: value}
	list.Len++

	if list.Head == nil {
		list.Head, list.Tail = node, node
		return
	}

	list.Tail, list.Tail.Next = node, node
}

func reorderLogFiles(logs []string) []string {
	digits, letters := &LinkedList{}, &LinkedList{}
	for i := 0; i < len(logs); i++ {
		index := getLetterLogContentIndex(logs[i])
		if index >= 0 {
			letters.InsertAndSort(logs[i], index)
		} else {
			digits.Insert(logs[i])
		}
	}

	return toArray(letters, digits)
}

func toArray(first, second *LinkedList) []string {
	result := make([]string, 0, first.Len+second.Len)
	for first.Head != nil {
		result = append(result, first.Head.Value)
		first.Head = first.Head.Next
	}

	for second.Head != nil {
		result = append(result, second.Head.Value)
		second.Head = second.Head.Next
	}

	return result
}

func getLetterLogContentIndex(log string) int {
	i := 0
	for ; log[i] != ' '; i++ {
	}
	i++

	if 'a' <= log[i] && log[i] <= 'z' {
		return i
	}

	return -1
}

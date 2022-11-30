package main

type ImmutableListNode interface {
	getNext() ImmutableListNode
	printValue()
}

func printLinkedListInReverse(head ImmutableListNode) {
	if head == nil {
		return
	}

	printLinkedListInReverse(head.getNext())
	head.printValue()
}

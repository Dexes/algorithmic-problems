package main

type MyCalendarTwo struct {
	root *Node
}

type Node struct {
	val    int
	update int
	start  int
	end    int
	left   *Node
	right  *Node
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{root: &Node{start: 0, end: 1e9}}
}

func (c MyCalendarTwo) Book(start int, end int) bool {
	if getMax(c.root, start, end) > 1 {
		return false
	}

	increment(c.root, start, end)
	return true
}

func getMax(root *Node, start, end int) int {
	if start >= end {
		return 0
	}

	if root.left == nil {
		middle := (root.start + root.end) / 2
		root.left = &Node{start: root.start, end: middle}
		root.right = &Node{start: middle, end: root.end}
	}

	if root.update > 0 {
		root.left.update += root.update
		root.right.update += root.update

		root.val += root.update
		root.update = 0
	}

	if root.start == start && root.end == end {
		return root.val
	}

	return max(
		getMax(root.left, start, min(end, root.left.end)),
		getMax(root.right, max(start, root.right.start), end),
	)
}

func increment(root *Node, start, end int) {
	if start >= end {
		return
	}

	if root.start == start && root.end == end {
		root.val++
		root.left.update++
		root.right.update++
		return
	}

	increment(root.left, start, min(end, root.left.end))
	increment(root.right, max(start, root.right.start), end)
	root.val = max(root.left.val, root.right.val)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

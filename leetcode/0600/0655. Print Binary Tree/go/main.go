package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	result := makeMatrix(root)
	printNode(result, root, 1, len(result[0])/2)

	return result
}

func printNode(view [][]string, node *TreeNode, depth, middle int) {
	if node == nil {
		return
	}

	view[depth-1][middle] = toString(node.Val)
	if depth == len(view) {
		return
	}

	shift := 1 << (len(view) - depth - 1)
	printNode(view, node.Left, depth+1, middle-shift)
	printNode(view, node.Right, depth+1, middle+shift)
}

func makeMatrix(root *TreeNode) [][]string {
	depth := maxDepth(root)
	width := (1 << depth) - 1

	result := make([][]string, depth)
	for i := 0; i < depth; i++ {
		result[i] = make([]string, width)
	}

	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if left, right := maxDepth(root.Left), maxDepth(root.Right); left > right {
		return left + 1
	} else {
		return right + 1
	}
}

var buffer = []byte{0, 0, 0}

func toString(x int) string {
	isNegative, i := x < 0, 0
	if isNegative {
		x, i, buffer[0] = -x, 1, '-'
	}

	if x < 10 {
		buffer[i] = byte(x + '0')
		return string(buffer[:i+1])
	}

	buffer[i] = byte(x/10 + '0')
	buffer[i+1] = byte(x%10 + '0')
	return string(buffer[:i+2])
}

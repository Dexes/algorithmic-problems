package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StackNode struct {
	Node  *TreeNode
	Depth int
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0, 100)
	stack, stackIndex := make([]StackNode, 100), 0
	stack[stackIndex].Node, stack[stackIndex].Depth = root, 0
	processedDepth := -1

	for stackIndex >= 0 {
		node, depth := stack[stackIndex].Node, stack[stackIndex].Depth
		stackIndex--

		for ; node != nil; node = node.Right {
			if depth > processedDepth {
				result = append(result, node.Val)
				processedDepth = depth
			}

			depth++

			if node.Left != nil {
				stackIndex++
				stack[stackIndex].Node, stack[stackIndex].Depth = node.Left, depth
			}
		}
	}

	return result
}

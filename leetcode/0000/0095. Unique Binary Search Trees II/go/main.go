package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var capacity = []int{1, 1, 2, 5, 14, 42, 132, 429, 1430}
var emptyResult = []*TreeNode{nil}

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(min, max int) []*TreeNode {
	if min > max {
		return emptyResult
	}

	if min == max {
		return []*TreeNode{{Val: min}}
	}

	result := make([]*TreeNode, 0, capacity[max-min+1])

	for i := min; i <= max; i++ {
		leftTree, rightTree := generate(min, i-1), generate(i+1, max)

		for _, left := range leftTree {
			for _, right := range rightTree {
				result = append(result, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}

	return result
}

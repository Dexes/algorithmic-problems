package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var (
		first, second               *TreeNode
		firstSuccess, secondSuccess bool
		firstChannel, secondChannel = make(chan *TreeNode, 10), make(chan *TreeNode, 10)
	)

	go getLeaves(root1, firstChannel)
	go getLeaves(root2, secondChannel)

	for {
		first, firstSuccess = <-firstChannel
		second, secondSuccess = <-secondChannel
		if !firstSuccess || !secondSuccess {
			return firstSuccess == secondSuccess
		}

		if first.Val != second.Val {
			return false
		}
	}
}

func getLeaves(root *TreeNode, writer chan<- *TreeNode) {
	dfs(root, writer)
	close(writer)
}

func dfs(root *TreeNode, writer chan<- *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		writer <- root
		return
	}

	dfs(root.Left, writer)
	dfs(root.Right, writer)
}

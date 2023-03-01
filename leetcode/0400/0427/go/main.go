package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return toNode(grid, 0, 0, len(grid), len(grid))
}

func toNode(grid [][]int, top, left, bottom, right int) *Node {
	node := &Node{Val: grid[top][left] == 1, IsLeaf: isLeaf(grid, top, left, bottom, right)}
	if !node.IsLeaf {
		bHalf, rHalf := (top+bottom)>>1, (left+right)>>1
		node.TopLeft = toNode(grid, top, left, bHalf, rHalf)
		node.TopRight = toNode(grid, top, rHalf, bHalf, right)
		node.BottomLeft = toNode(grid, bHalf, left, bottom, rHalf)
		node.BottomRight = toNode(grid, bHalf, rHalf, bottom, right)
	}

	return node
}

func isLeaf(grid [][]int, top, left, bottom, right int) bool {
	for i := top; i < bottom; i++ {
		for j := left; j < right; j++ {
			if grid[i][j] != grid[top][left] {
				return false
			}
		}
	}

	return true
}

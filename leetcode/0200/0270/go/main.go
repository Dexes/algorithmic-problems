package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	_, result := search(root, target, abs(float64(root.Val)-target), root.Val)
	return result
}

func search(root *TreeNode, target, diff float64, closest int) (float64, int) {
	if root == nil {
		return diff, closest
	}

	currentDiff := float64(root.Val) - target
	if x := abs(currentDiff); x < diff {
		diff, closest = x, root.Val
	}

	if currentDiff > 0 {
		diff, closest = search(root.Left, target, diff, closest)
	} else {
		diff, closest = search(root.Right, target, diff, closest)
	}

	return diff, closest
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}

	return a
}

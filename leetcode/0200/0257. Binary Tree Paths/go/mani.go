package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	return dfs(root, make([]int, 0, 100), make([]string, 0, 100))
}

func dfs(node *TreeNode, path []int, paths []string) []string {
	path = append(path, node.Val)

	if node.Left == nil && node.Right == nil {
		return append(paths, toString(path))
	}

	if node.Left != nil {
		paths = dfs(node.Left, path, paths)
	}

	if node.Right != nil {
		paths = dfs(node.Right, path, paths)
	}

	return paths
}

func toString(path []int) string {
	result := make([]byte, 0, 100)

	for i := len(path) - 1; i >= 0; i-- {
		num, isNegative := path[i], false
		if num < 0 {
			num, isNegative = -num, true
		}

		for num > 0 {
			result = append(result, byte(num%10)+'0')
			num /= 10
		}

		if isNegative {
			result = append(result, '-')
		}

		result = append(result, '>', '-')
	}

	if len(result) > 100 {
		return strconv.Itoa(len(result))
	}

	return string(reverse(result[:len(result)-2]))
}

func reverse(bytes []byte) []byte {
	for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}

	return bytes
}

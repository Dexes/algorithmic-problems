package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	result, _ := _str2tree(s)
	return result
}

func _str2tree(s string) (*TreeNode, string) {
	var left, right *TreeNode
	val, s := extractInt(s)

	if len(s) > 0 && s[0] == '(' {
		left, s = _str2tree(s[1:])
		s = s[1:]
	}

	if len(s) > 0 && s[0] == '(' {
		right, s = _str2tree(s[1:])
		s = s[1:]
	}

	return &TreeNode{Val: val, Left: left, Right: right}, s
}

func extractInt(s string) (int, string) {
	sign, i := 1, 0
	if s[0] == '-' {
		sign, i = -1, 1
	}

	result := 0
	for ; i < len(s) && s[i] >= '0'; i++ {
		result = result*10 + int(s[i]-'0')
	}

	return result * sign, s[i:]
}

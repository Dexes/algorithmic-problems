package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const difference = 1001

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	return string(write(make([]byte, 0, 10000), root))
}

func (c *Codec) deserialize(data string) *TreeNode {
	result, _ := read(data, 0)
	return result
}

func write(buffer []byte, root *TreeNode) []byte {
	if root == nil {
		return append(buffer, 0, 0)
	}

	value := root.Val + difference
	buffer = append(buffer, byte(value>>8), byte(value))

	return write(write(buffer, root.Left), root.Right)
}

func read(s string, index int) (*TreeNode, int) {
	value := int(s[index])<<8 + int(s[index+1])
	if value == 0 {
		return nil, 2
	}

	left, leftSize := read(s, index+2)
	right, rightSize := read(s, index+2+leftSize)

	return &TreeNode{Val: value - difference, Left: left, Right: right}, 2 + leftSize + rightSize
}

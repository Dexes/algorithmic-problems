package main

type NumArray struct {
	tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)*2)
	for i := len(nums); i < len(tree); i++ {
		tree[i] = nums[i-len(nums)]
	}

	for i := len(nums) - 1; i > 0; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}

	return NumArray{tree: tree}
}

func (a *NumArray) Update(index int, val int) {
	tree := a.tree
	index += len(tree) / 2
	tree[index] = val

	for index /= 2; index > 0; index /= 2 {
		a.tree[index] = a.tree[index*2] + a.tree[index*2+1]
	}
}

func (a *NumArray) SumRange(left int, right int) (result int) {
	tree := a.tree
	left, right = left+len(tree)/2, right+len(tree)/2+1

	for left < right {
		if left&1 == 1 {
			result += tree[left]
			left++
		}

		if right&1 == 1 {
			right--
			result += tree[right]
		}

		left /= 2
		right /= 2
	}

	return result
}

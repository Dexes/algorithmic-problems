package main

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return NumArray{prefixSum: nums}
}

func (x *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return x.prefixSum[right]
	}

	return x.prefixSum[right] - x.prefixSum[left-1]
}

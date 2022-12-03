package main

func numberOfArithmeticSlices(nums []int) int {
	stopIndex, lastIndex, j, diff := len(nums)-3, len(nums)-1, 0, 0
	result := 0

	for i := 0; i <= stopIndex; i += j {
		for j, diff = i+1, nums[i]-nums[i+1]; j < lastIndex && nums[j]-nums[j+1] == diff; j++ {
		}

		if j -= i; j > 1 {
			result += j * (j - 1) / 2
		}
	}

	return result
}

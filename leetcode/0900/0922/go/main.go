package main

func sortArrayByParityII(nums []int) []int {
	evenIndex, oddIndex := 0, 1
	for {
		for ; evenIndex < len(nums) && nums[evenIndex]&1 == 0; evenIndex += 2 {
		}

		for ; oddIndex < len(nums) && nums[oddIndex]&1 == 1; oddIndex += 2 {
		}

		if evenIndex == len(nums) || oddIndex == len(nums) {
			break
		}

		nums[evenIndex], nums[oddIndex] = nums[oddIndex], nums[evenIndex]
		evenIndex, oddIndex = evenIndex+2, oddIndex+2
	}

	return nums
}

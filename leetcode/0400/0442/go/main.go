package main

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		rearrange(nums, i)
	}

	rIndex, wIndex := 0, 0
	for ; rIndex < len(nums); rIndex++ {
		if num := rIndex + 1; nums[rIndex] < -num {
			nums[wIndex], wIndex = num, wIndex+1
		}
	}

	return nums[:wIndex]
}

func rearrange(nums []int, index int) {
	if nums[index] < 1 {
		return
	}

	for {
		target := nums[index] - 1
		if target == index {
			nums[index] = -nums[index]
			return
		}

		if nums[target] <= 0 {
			nums[index], nums[target] = 0, nums[target]-nums[index]
			return
		}

		nums[index], nums[target] = nums[target], -nums[index]
	}
}

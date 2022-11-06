package main

func applyOperations(nums []int) []int {
	rIndex, wIndex := 0, 0
	for ; rIndex < len(nums); rIndex++ {
		if nums[rIndex] == 0 {
			continue
		}

		if next := rIndex + 1; next < len(nums) && nums[rIndex] == nums[next] {
			nums[wIndex] = nums[rIndex] << 1
			rIndex, wIndex = rIndex+1, wIndex+1
			continue
		}

		nums[wIndex] = nums[rIndex]
		wIndex++
	}

	for ; wIndex < len(nums); wIndex++ {
		nums[wIndex] = 0
	}

	return nums
}

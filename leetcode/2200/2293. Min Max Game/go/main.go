package main

func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		rIndex, wIndex, even := 0, 0, true
		for rIndex < len(nums) {
			if x, y := nums[rIndex], nums[rIndex+1]; (even && x < y) || (!even && x > y) {
				nums[wIndex] = x
			} else {
				nums[wIndex] = y
			}

			rIndex, wIndex, even = rIndex+2, wIndex+1, !even
		}

		nums = nums[:wIndex]
	}

	return nums[0]
}

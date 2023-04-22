package main

func productExceptSelf(nums []int) []int {
	product, zeroPosition := 1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zeroPosition >= 0 {
				return toZeros(nums)
			}

			zeroPosition = i
			continue
		}

		product *= nums[i]
	}

	if zeroPosition >= 0 {
		nums = toZeros(nums)
		nums[zeroPosition] = product
		return nums
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = product / nums[i]
	}

	return nums
}

func toZeros(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

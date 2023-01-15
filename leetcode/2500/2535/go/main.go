package main

func differenceOfSum(nums []int) (result int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] < 10 {
			continue
		}

		result += nums[i] - digits(nums[i])
	}

	return result
}

func digits(x int) (result int) {
	for ; x > 0; x /= 10 {
		result += x % 10
	}

	return result
}

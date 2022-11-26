package main

func sumOfDigits(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	sum := 0
	for ; min > 0; min /= 10 {
		sum += min % 10
	}

	return (sum + 1) & 1
}

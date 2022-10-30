package main

func averageValue(nums []int) int {
	sum, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%6 == 0 {
			sum, count = sum+nums[i], count+1
		}
	}

	if count == 0 {
		return 0
	}

	return sum / count
}

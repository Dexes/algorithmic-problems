package main

func maxProduct(nums []int) int {
	first, second := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > second {
			first, second = second, nums[i]
			continue
		}

		if nums[i] > first {
			first = nums[i]
		}
	}

	return (first - 1) * (second - 1)
}

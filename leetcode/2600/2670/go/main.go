package main

const max = 51

func distinctDifferenceArray(nums []int) []int {
	left, right := [max]int{}, [max]int{}
	lCount, rCount := 0, 0

	for i := 0; i < len(nums); i++ {
		if right[nums[i]]++; right[nums[i]] == 1 {
			rCount++
		}
	}

	for i := 0; i < len(nums); i++ {
		if left[nums[i]]++; left[nums[i]] == 1 {
			lCount++
		}

		if right[nums[i]]--; right[nums[i]] == 0 {
			rCount--
		}

		nums[i] = lCount - rCount
	}

	return nums
}

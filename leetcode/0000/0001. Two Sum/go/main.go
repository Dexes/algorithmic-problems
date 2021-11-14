package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		x := target - nums[i]

		for j := i + 1; j < len(nums); j++ {
			if nums[j] == x {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

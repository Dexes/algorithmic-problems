package main

func countQuadruplets(nums []int) int {
	result, count := 0, make([]int, 201)

	for second := 0; second < len(nums); second++ {
		for first := 0; first < second; first++ {
			count[nums[first]+nums[second]]++
		}

		for third, sumIndex := second+1, second+2; sumIndex < len(nums); sumIndex++ {
			if nums[sumIndex]-nums[third] > 0 {
				result += count[nums[sumIndex]-nums[third]]
			}
		}
	}

	return result
}

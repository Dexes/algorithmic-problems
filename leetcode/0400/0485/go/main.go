package main

func findMaxConsecutiveOnes(nums []int) int {
	result, number := 0, 0
	for i := 0; i < len(nums); i += number {
		i += countConsecutive(nums, 0, i)
		number = countConsecutive(nums, 1, i)
		if number > result {
			result = number
		}
	}

	return result
}

func countConsecutive(nums []int, x, start int) int {
	i := start
	for ; i < len(nums) && nums[i] == x; i++ {
	}

	return i - start
}

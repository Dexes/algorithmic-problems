package main

func prefixesDivBy5(nums []int) []bool {
	result, number := make([]bool, len(nums)), 0
	for i := 0; i < len(nums); i++ {
		number = (number<<1 | nums[i]) % 5
		result[i] = number == 0
	}

	return result
}

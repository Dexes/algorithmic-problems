package main

func findThePrefixCommonArray(a, b []int) []int {
	nums, numsLength := make([]bool, len(a)+1), 0
	for i := 0; i < len(a); i++ {
		numsLength += invert(nums, a[i])
		numsLength += invert(nums, b[i])

		a[i] = i + 1 - (numsLength / 2)
	}

	return a
}

func invert(nums []bool, num int) int {
	if nums[num] = !nums[num]; nums[num] {
		return 1
	}

	return -1
}

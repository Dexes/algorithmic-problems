package main

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	a, b, c := nums[0], nums[1], nums[2]
	a, b = sort(a, b)
	a, c = sort(a, c)
	b, c = sort(b, c)

	return b
}

func sort(a, b int) (int, int) {
	if a > b {
		return a, b
	}

	return b, a
}

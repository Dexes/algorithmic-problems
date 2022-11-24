package main

func lexicalOrder(n int) []int {
	result := make([]int, 0, n)
	for i := 1; i < 10 && i <= n; i++ {
		result = append(result, i)
		result = fill(n, i, result)
	}

	return result
}

func fill(max int, num int, nums []int) []int {
	var current int
	num *= 10

	for i := 0; i < 10; i++ {
		if current = num + i; current > max {
			return nums
		}

		nums = append(nums, current)
		nums = fill(max, current, nums)
	}

	return nums
}

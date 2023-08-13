package main

func maxSum(nums []int) int {
	var max [10]int
	result := -1

	for _, x := range nums {
		d := getMaxDigit(x)
		if max[d] > 0 && max[d]+x > result {
			result = max[d] + x
		}

		if x > max[d] {
			max[d] = x
		}
	}

	return result
}

func getMaxDigit(x int) (result int) {
	for ; x > 0; x /= 10 {
		if y := x % 10; y > result {
			result = y
		}
	}

	return result
}

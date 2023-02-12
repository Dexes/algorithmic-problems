package main

func findTheArrayConcVal(nums []int) int64 {
	result := 0
	l, r := 0, len(nums)-1
	for ; l < r; l, r = l+1, r-1 {
		result += concat(nums[l], nums[r])
	}

	if l == r {
		return int64(result + nums[l])
	}

	return int64(result)
}

func concat(a, b int) int {
	for x := b; x > 0; x /= 10 {
		a *= 10
	}

	return a + b
}

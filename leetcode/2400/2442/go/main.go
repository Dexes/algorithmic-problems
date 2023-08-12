package main

const max = 1e6 + 1

func countDistinctIntegers(nums []int) (result int) {
	var set [max]bool

	for _, x := range nums {
		if !set[x] {
			set[x] = true
			result++
		}

		if y := reverse(x); !set[y] {
			set[y] = true
			result++
		}
	}

	return result
}

func reverse(x int) (y int) {
	for ; x > 0; x /= 10 {
		y = y*10 + x%10
	}

	return y
}

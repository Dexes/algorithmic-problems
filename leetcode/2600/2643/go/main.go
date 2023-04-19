package main

func rowAndMaximumOnes(mat [][]int) []int {
	row, max := 0, 0
	for i := 0; i < len(mat); i++ {
		if x := sum(mat[i]); x > max {
			row, max = i, x
		}
	}

	return []int{row, max}
}

func sum(nums []int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}

	return result
}

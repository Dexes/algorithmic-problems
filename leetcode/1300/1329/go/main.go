package main

import (
	"sort"
)

func diagonalSort(mat [][]int) [][]int {
	nums := make([]int, 0, min(len(mat), len(mat[0])))
	for diagonal := 1 - len(mat); diagonal < len(mat[0]); diagonal++ {
		sortDiagonal(mat, diagonal, nums)
	}

	return mat
}

func sortDiagonal(mat [][]int, diagonal int, nums []int) {
	row, col := max(0, -diagonal), max(0, diagonal)
	for r, c := row, col; r < len(mat) && c < len(mat[0]); r, c = r+1, c+1 {
		nums = append(nums, mat[r][c])
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		mat[row+i][col+i] = nums[i]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

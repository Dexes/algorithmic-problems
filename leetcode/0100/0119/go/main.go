package main

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i >> 1; j > 0; j-- {
			result[j] = result[j] + result[j-1]
		}

		if i&1 == 1 {
			result[i>>1+1] = result[i>>1]
		}
	}

	for i := rowIndex>>1 + 1; i <= rowIndex; i++ {
		result[i] = result[rowIndex-i]
	}

	return result
}

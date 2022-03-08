package main

func numSpecial(mat [][]int) int {
	rSums, cSums := getSums(mat)
	result := 0

	for rNumber := 0; rNumber < len(rSums); rNumber++ {
		if rSums[rNumber] != 1 {
			continue
		}

		for cNumber := 0; cNumber < len(cSums); cNumber++ {
			if mat[rNumber][cNumber] == 1 && cSums[cNumber] == 1 {
				result++
			}
		}
	}

	return result
}

func getSums(mat [][]int) ([]int, []int) {
	rSums, cSums := make([]int, len(mat)), make([]int, len(mat[0]))
	for rNumber := 0; rNumber < len(rSums); rNumber++ {
		for cNumber := 0; cNumber < len(cSums); cNumber++ {
			rSums[rNumber] += mat[rNumber][cNumber]
			cSums[cNumber] += mat[rNumber][cNumber]
		}
	}

	return rSums, cSums
}

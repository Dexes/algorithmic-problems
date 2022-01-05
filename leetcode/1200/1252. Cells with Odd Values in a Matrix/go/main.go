package main

func oddCells(m int, n int, indices [][]int) int {
	rSums, cSums := make([]int, m), make([]int, n)
	for i := 0; i < len(indices); i++ {
		rSums[indices[i][0]]++
		cSums[indices[i][1]]++
	}

	re, ro := countEvenAndOdd(rSums)
	ce, co := countEvenAndOdd(cSums)

	return re*co + ce*ro
}

func countEvenAndOdd(nums []int) (int, int) {
	even, odd := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			odd++
		} else {
			even++
		}
	}

	return even, odd
}

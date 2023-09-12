package main

func numberOfPoints(nums [][]int) (result int) {
	var points [102]int8
	for _, rng := range nums {
		points[rng[0]]++
		points[rng[1]+1]--
	}

	cars := int8(0)
	for _, x := range points {
		if cars += x; cars > 0 {
			result++
		}
	}

	return result
}

package main

func groupThePeople(sizes []int) [][]int {
	groups := make([][]int, max(sizes)+1)
	result := make([][]int, 0, len(sizes)/4)

	for i, size := range sizes {
		if groups[size] == nil {
			groups[size] = make([]int, 0, size)
		}

		groups[size] = append(groups[size], i)
		if len(groups[size]) == size {
			result, groups[size] = append(result, groups[size]), nil
		}
	}

	return result
}

func max(nums []int) (result int) {
	for _, x := range nums {
		if x > result {
			result = x
		}
	}

	return result
}

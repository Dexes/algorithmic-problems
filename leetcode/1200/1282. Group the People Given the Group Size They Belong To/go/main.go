package main

func groupThePeople(groupSizes []int) [][]int {
	var group []int
	groups := make(map[int][]int)
	result := make([][]int, 0, len(groupSizes)/4)

	for i, size := range groupSizes {
		if group = groups[size]; len(group) == 0 {
			group = make([]int, 0, size)
		}

		group = append(group, i)
		if len(group) == size {
			result = append(result, group)
			group = nil
		}

		groups[size] = group
	}

	return result
}

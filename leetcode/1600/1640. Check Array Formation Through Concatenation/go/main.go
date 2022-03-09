package main

func canFormArray(arr []int, pieces [][]int) bool {
	indices := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		indices[arr[i]] = i
	}

	length, prevIndex, currentIndex, ok := 0, 0, 0, true
	for i := 0; i < len(pieces); i++ {
		if prevIndex, ok = indices[pieces[i][0]]; !ok {
			return false
		}

		for j := 1; j < len(pieces[i]); j++ {
			if currentIndex, ok = indices[pieces[i][j]]; !ok || prevIndex != currentIndex-1 {
				return false
			}

			prevIndex = currentIndex
		}

		length += len(pieces[i])
	}

	return len(arr) == length
}

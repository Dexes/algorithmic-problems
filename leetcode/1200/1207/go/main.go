package main

func uniqueOccurrences(arr []int) bool {
	x := make(map[int]struct{})
	for _, number := range count(arr) {
		if _, ok := x[number]; ok {
			return false
		}

		x[number] = struct{}{}
	}

	return true
}

func count(arr []int) map[int]int {
	result := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		result[arr[i]]++
	}

	return result
}

package main

func containsPattern(arr []int, length, repeats int) bool {
	for i, stopIndex := 0, len(arr)-length*repeats; i <= stopIndex; i++ {
		if check(arr, length, repeats, i) {
			return true
		}
	}

	return false
}

func check(arr []int, length, repeats, index int) bool {
	for i := 0; i < length; i++ {
		firstIndex, secondIndex, stopIndex := index+i, index+i+length, index+i+length*repeats
		for ; secondIndex < stopIndex; secondIndex += length {
			if arr[firstIndex] != arr[secondIndex] {
				return false
			}
		}
	}

	return true
}

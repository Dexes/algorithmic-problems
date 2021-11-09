package main

func findMedianSortedArrays(first, second []int) float64 {
	firstLength, secondLength := len(first), len(second)
	firstIndex, secondIndex := 0, 0
	length := firstLength + secondLength
	medianIndex := length / 2
	even := length|1 != length
	prev, current := 0, 0

	for firstIndex+secondIndex <= medianIndex {
		if firstIndex >= firstLength {
			prev, current = current, second[secondIndex]
			secondIndex++
			continue
		}

		if secondIndex >= secondLength || first[firstIndex] < second[secondIndex] {
			prev, current = current, first[firstIndex]
			firstIndex++
			continue
		}

		prev, current = current, second[secondIndex]
		secondIndex++
	}

	if even {
		return float64(prev+current) / 2
	}

	return float64(current)
}

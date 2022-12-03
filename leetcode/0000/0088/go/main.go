package main

func merge(first []int, firstIndex int, second []int, secondIndex int) {
	firstIndex--
	secondIndex--

	for firstIndex >= 0 || secondIndex >= 0 {
		if firstIndex < 0 {
			first[firstIndex+secondIndex+1] = second[secondIndex]
			secondIndex--
			continue
		}

		if secondIndex < 0 || first[firstIndex] > second[secondIndex] {
			first[firstIndex+secondIndex+1] = first[firstIndex]
			firstIndex--
			continue
		}

		first[firstIndex+secondIndex+1] = second[secondIndex]
		secondIndex--
	}
}

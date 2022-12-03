package main

func arrayStringsAreEqual(first []string, second []string) bool {
	firstIndex, secondIndex := 0, 0
	firstElementIndex, secondElementIndex := 0, 0
	firstEnded, secondEnded := false, false

	for {
		if firstEnded && secondEnded {
			return true
		}

		if firstEnded != secondEnded {
			return false
		}

		if first[firstIndex][firstElementIndex] != second[secondIndex][secondElementIndex] {
			return false
		}

		firstIndex, firstElementIndex, firstEnded = nextIndices(first, firstIndex, firstElementIndex)
		secondIndex, secondElementIndex, secondEnded = nextIndices(second, secondIndex, secondElementIndex)
	}
}

func nextIndices(data []string, index, elementIndex int) (int, int, bool) {
	if elementIndex+1 < len(data[index]) {
		return index, elementIndex + 1, false
	}

	if index+1 < len(data) {
		return index + 1, 0, false
	}

	return 0, 0, true
}

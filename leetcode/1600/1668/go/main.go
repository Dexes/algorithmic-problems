package main

func maxRepeating(sequence string, word string) int {
	result, stopIndex := 0, len(sequence)-len(word)
	for i := 0; i <= stopIndex; i++ {
		if count := countRepeats(sequence, word, i); count > result {
			result, stopIndex = count, len(sequence)-len(word)*(count+1)
		}
	}

	return result
}

func countRepeats(sequence, word string, index int) int {
	result, stop := 0, len(sequence)-len(word)
	for ; index <= stop; index += len(word) {
		for i := 0; i < len(word); i++ {
			if sequence[i+index] != word[i] {
				return result
			}
		}

		result++
	}

	return result
}

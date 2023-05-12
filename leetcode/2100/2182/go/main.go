package main

func repeatLimitedString(s string, repeatLimit int) string {
	letters := make([]int, 123)
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
	}

	result := make([]byte, 0, len(s))
	letter, number := getLetter(letters, 'z')
	nextLetter, nextNumber := getLetter(letters, letter-1)
	for number > 0 {
		result = insert(result, letter, min(number, repeatLimit))
		number -= repeatLimit

		if nextNumber <= 0 {
			return string(result)
		}

		if number <= 0 {
			letter, number = nextLetter, nextNumber
			nextLetter, nextNumber = getLetter(letters, nextLetter-1)
			continue
		}

		result = append(result, nextLetter)
		if nextNumber--; nextNumber == 0 {
			nextLetter, nextNumber = getLetter(letters, nextLetter-1)
		}
	}

	return string(result)
}

func insert(bytes []byte, letter byte, number int) []byte {
	for ; number > 0; number-- {
		bytes = append(bytes, letter)
	}

	return bytes
}

func getLetter(letters []int, from byte) (byte, int) {
	for ; from < 255; from-- {
		if letters[from] > 0 {
			return from, letters[from]
		}
	}

	return 0, 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

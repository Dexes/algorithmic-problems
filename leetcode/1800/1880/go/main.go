package main

var base = 'j' - 'a' + byte(1)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	targetWord = trim(targetWord)

	i, j, k := len(firstWord)-1, len(secondWord)-1, len(targetWord)-1
	s, b := byte(0), byte(0)

	for i >= 0 || j >= 0 {
		s, b = sum(b, digit(firstWord, i), digit(secondWord, j))
		if s != digit(targetWord, k) {
			return false
		}

		i, j, k = i-1, j-1, k-1
	}

	return k < 0 && b == 0 || k == 0 && digit(targetWord, k) == b
}

func trim(s string) string {
	i := 0
	for ; i < len(s) && s[i] == 'a'; i++ {
	}

	return s[i:]
}

func digit(s string, index int) byte {
	if index < 0 {
		return 0
	}

	return s[index] - 'a'
}

func sum(b, x, y byte) (byte, byte) {
	result := b + x + y
	if result >= base {
		return result - base, 1
	}

	return result, 0
}

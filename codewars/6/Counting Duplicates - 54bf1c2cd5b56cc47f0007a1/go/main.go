package kata

func duplicate_count(s string) int {
	counter, letter := make([]int, 91), byte(0)
	for i := 0; i < len(s); i++ {
		if letter = s[i]; letter > 'Z' {
			letter -= 32
		}

		counter[letter]++
	}

	return count(counter, 'A', 'Z') + count(counter, '0', '9')
}

func count(counter []int, a, b byte) (result int) {
	for ; a <= b; a++ {
		if counter[a] > 1 {
			result++
		}
	}

	return result
}

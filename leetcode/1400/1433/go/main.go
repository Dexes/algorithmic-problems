package main

func checkIfCanBreak(s1 string, s2 string) bool {
	if len(s1) == 1 {
		return true
	}

	c1, c2 := [123]int{}, [123]int{}
	for i := 0; i < len(s1); i++ {
		c1[s1[i]]++
		c2[s2[i]]++
	}

	diff, isNegative, isEmpty := 0, false, true
	for i := 'a'; i <= 'z'; i++ {
		if c1[i] == c2[i] {
			continue
		}

		if isEmpty {
			diff, isNegative, isEmpty = c1[i]-c2[i], c1[i] < c2[i], false
			continue
		}

		diff += c1[i] - c2[i]
		if diff != 0 && (diff < 0 != isNegative) {
			return false
		}
	}

	return true
}

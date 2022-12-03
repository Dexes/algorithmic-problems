package main

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	index, counter := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}

		switch counter {
		case 0:
			index, counter = i, 1
		case 1:
			if s1[index] != s2[i] || s2[index] != s1[i] {
				return false
			}

			counter = 2
		default:
			return false
		}
	}

	return counter != 1
}

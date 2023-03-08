package main

func shortestWay(source, target string) (result int) {
	if s, t := letters(source), letters(target); (s & t) < t {
		return -1
	}

	for tIndex := 0; tIndex < len(target); {
		for sIndex := 0; tIndex < len(target) && sIndex < len(source); sIndex++ {
			if source[sIndex] == target[tIndex] {
				tIndex++
			}
		}

		result++
	}

	return result
}

func letters(s string) (result int) {
	for i := 0; i < len(s); i++ {
		result |= 1 << (s[i] - 'a')
	}

	return result
}

package main

func checkAlmostEquivalent(first string, second string) bool {
	counter := make([]int, 26)
	for i := 0; i < len(first); i++ {
		counter[first[i]-'a']++
		counter[second[i]-'a']--
	}

	for i := 0; i < len(counter); i++ {
		if counter[i] < -3 || counter[i] > 3 {
			return false
		}
	}

	return true
}

package main

func percentageLetter(s string, letter byte) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			counter += 100
		}
	}

	return counter / len(s)
}

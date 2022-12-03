package main

func checkRecord(s string) bool {
	aCount, lLength := 0, 0
	for i := 0; i < len(s) && lLength < 3; i++ {
		if s[i] == 'L' {
			lLength++
			continue
		}

		lLength = 0
		if s[i] == 'A' {
			aCount++
		}
	}

	return aCount < 2 && lLength < 3
}

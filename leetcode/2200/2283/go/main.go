package main

func digitCount(num string) bool {
	counter := make([]byte, 10)
	for i := 0; i < len(num); i++ {
		counter[num[i]-'0']++
	}

	for i := 0; i < len(num); i++ {
		if num[i]-'0' != counter[i] {
			return false
		}
	}

	return true
}

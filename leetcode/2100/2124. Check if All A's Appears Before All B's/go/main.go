package main

func checkString(s string) bool {
	i := 0
	for ; i < len(s) && s[i] == 'a'; i++ {
	}

	for ; i < len(s) && s[i] == 'b'; i++ {
	}

	return i == len(s)
}

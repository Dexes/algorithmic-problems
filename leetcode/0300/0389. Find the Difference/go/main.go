package main

func findTheDifference(s string, t string) byte {
	result := byte(0)
	for i := 0; i < len(s); i++ {
		result ^= s[i] ^ t[i]
	}

	return result ^ t[len(t)-1]
}

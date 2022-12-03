package main

func findArray(prefix []int) []int {
	x := prefix[0]
	for i := 1; i < len(prefix); i++ {
		prefix[i], x = prefix[i]^x, prefix[i]
	}

	return prefix
}

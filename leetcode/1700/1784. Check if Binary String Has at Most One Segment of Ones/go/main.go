package main

func checkOnesSegment(s string) bool {
	index := len(s) - 1

	for ; s[index] == '0'; index-- {
	}

	for ; index > 0 && s[index] == '1'; index-- {
	}

	return index == 0
}

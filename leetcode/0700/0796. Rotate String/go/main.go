package main

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	for shift := -1; ; {
		shift = search(goal, s[0], shift+1)
		if shift == -1 {
			return false
		}

		if check(s, goal, shift) {
			return true
		}
	}
}

func check(s string, goal string, shift int) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != goal[(i+shift)%len(goal)] {
			return false
		}
	}

	return true
}

func search(s string, letter byte, start int) int {
	for i := start; i < len(s); i++ {
		if s[i] == letter {
			return i
		}
	}

	return -1
}

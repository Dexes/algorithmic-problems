package main

func isLongPressedName(name string, typed string) bool {
	i, j, x, y := 0, 0, 0, 0
	for ; i < len(name) && j < len(typed); i, j = i+x, j+y {
		x, y = countChars(name, i), countChars(typed, j)
		if name[i] != typed[j] || x > y {
			return false
		}
	}

	return i == len(name) && j == len(typed)
}

func countChars(s string, index int) int {
	i := index + 1
	for ; i < len(s) && s[i] == s[index]; i++ {
	}

	return i - index
}

package main

func compareVersion(version1 string, version2 string) int {
	x, y, i, j := 0, 0, -1, -1

	for i < len(version1) || j < len(version2) {
		x, i = readInt(version1, i+1)
		y, j = readInt(version2, j+1)

		switch {
		case x < y:
			return -1
		case y < x:
			return 1
		}
	}

	return 0
}

func readInt(s string, index int) (int, int) {
	result := 0
	for ; index < len(s) && s[index] != '.'; index++ {
		result = result*10 + int(s[index]-'0')
	}

	return result, index
}

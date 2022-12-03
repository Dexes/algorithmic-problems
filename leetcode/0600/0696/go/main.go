package main

func countBinarySubstrings(s string) int {
	result, prev, current, bit := 0, 0, 0, byte('0')
	for i := 0; i < len(s); i += current {
		current = count(s, bit, i)
		result += min(prev, current)

		prev, bit = current, invert(bit)
	}

	return result
}

func count(s string, b byte, start int) int {
	i := start
	for ; i < len(s) && s[i] == b; i++ {
	}

	return i - start
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func invert(b byte) byte {
	if b == '0' {
		return '1'
	}

	return '0'
}

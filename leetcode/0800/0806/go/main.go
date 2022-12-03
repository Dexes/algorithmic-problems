package main

func numberOfLines(widths []int, s string) []int {
	lines, width := 0, 0
	for i := 0; i < len(s); i, width = write(widths, s, i) {
		lines++
	}

	return []int{lines, width}
}

func write(widths []int, s string, index int) (int, int) {
	width, current := 0, 0
	for ; index < len(s); index++ {
		current = widths[s[index]-'a']
		width += current
		if width > 100 {
			return index, width - current
		}
	}

	return index, width
}

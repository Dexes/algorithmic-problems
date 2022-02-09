package main

func secondHighest(s string) int {
	first, second := byte(0), byte(0)
	for i := 0; i < len(s); i++ {
		if first < s[i] && s[i] <= '9' {
			first, second = s[i], first
			continue
		}

		if second < s[i] && s[i] < first {
			second = s[i]
		}
	}

	if second == 0 {
		return -1
	}

	return int(second - '0')
}

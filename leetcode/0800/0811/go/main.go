package main

func subdomainVisits(domains []string) []string {
	stat := make(map[string]int)
	for i := 0; i < len(domains); i++ {
		for x, j := readInt(domains[i]); j < len(domains[i]); j = findDot(domains[i], j) {
			stat[domains[i][j:]] += x
		}
	}

	result, buffer, i := make([]string, 0, len(stat)), [10]byte{}, 0
	buffer[9] = ' '

	for domain, num := range stat {
		for i = 9; num > 0; num /= 10 {
			i--
			buffer[i] = byte(num%10) + '0'
		}

		result = append(result, string(buffer[i:])+domain)
	}

	return result
}

func findDot(s string, i int) int {
	for ; i < len(s) && s[i] != '.'; i++ {
	}

	return i + 1
}

func readInt(s string) (int, int) {
	result, index := 0, 0
	for ; index < len(s) && s[index] != ' '; index++ {
		result = result*10 + int(s[index]-'0')
	}

	return result, index + 1
}

package main

func validIPAddress(s string) string {
	switch {
	case isIPv4(s):
		return "IPv4"
	case isIPv6(s):
		return "IPv6"
	default:
		return "Neither"
	}
}

func isIPv4(s string) bool {
	if len(s) < 7 || 15 < len(s) {
		return false
	}

	indices, success := getIndices(s, '.', 3)
	if !success {
		return false
	}

	return checkIPv4Part(s[:indices[0]]) &&
		checkIPv4Part(s[indices[0]+1:indices[1]]) &&
		checkIPv4Part(s[indices[1]+1:indices[2]]) &&
		checkIPv4Part(s[indices[2]+1:])
}

func checkIPv4Part(s string) bool {
	switch len(s) {
	case 1:
		return '0' <= s[0] && s[0] <= '9'
	case 2:
		return '1' <= s[0] && s[0] <= '9' && '0' <= s[1] && s[1] <= '9'
	case 3:
		checkSum := int(s[0]-'0')*100 + int(s[1]-'0')*10 + int(s[2]-'0')
		return 100 <= checkSum && checkSum <= 255
	default:
		return false
	}
}

func isIPv6(s string) bool {
	if len(s) < 15 || 39 < len(s) {
		return false
	}

	indices, success := getIndices(s, ':', 7)
	if !success {
		return false
	}

	return checkIPv6Part(s[:indices[0]]) &&
		checkIPv6Part(s[indices[0]+1:indices[1]]) &&
		checkIPv6Part(s[indices[1]+1:indices[2]]) &&
		checkIPv6Part(s[indices[2]+1:indices[3]]) &&
		checkIPv6Part(s[indices[3]+1:indices[4]]) &&
		checkIPv6Part(s[indices[4]+1:indices[5]]) &&
		checkIPv6Part(s[indices[5]+1:indices[6]]) &&
		checkIPv6Part(s[indices[6]+1:])
}

func checkIPv6Part(s string) bool {
	switch len(s) {
	case 1:
		return isHex(s[0])
	case 2:
		return isHex(s[0]) && isHex(s[1])
	case 3:
		return isHex(s[0]) && isHex(s[1]) && isHex(s[2])
	case 4:
		return isHex(s[0]) && isHex(s[1]) && isHex(s[2]) && isHex(s[3])
	default:
		return false
	}
}

func isHex(x byte) bool {
	return '0' <= x && x <= '9' || 'a' <= x && x <= 'f' || 'A' <= x && x <= 'F'
}

func getIndices(s string, b byte, limit int) ([]int, bool) {
	result := make([]int, 0, limit)
	for i := 0; i < len(s); i++ {
		if s[i] != b {
			continue
		}

		if len(result) == limit {
			return nil, false
		}

		result = append(result, i)
	}

	return result, len(result) == limit
}

package kata

func Is_valid_ip(ip string) bool {
	index, length, isValid := 0, 0, false
	for i := 0; i < 4; i++ {
		if length, isValid = checkNumber(ip, index); !isValid {
			return false
		}

		index += length + 1
	}

	return true
}

func checkNumber(s string, index int) (length int, isValid bool) {
	switch i1, i2, i3 := index+1, index+2, index+3; {
	case index >= len(s):
		return -1, false

	case i1 == len(s) || s[i1] == '.':
		return 1, '0' <= s[index] && s[index] <= '9'

	case i2 == len(s) || s[i2] == '.':
		num := int(s[index]-'0')*10 + int(s[i1]-'0')
		return 2, s[index] > '0' && 0 <= num && num < 100

	case i3 == len(s) || s[i3] == '.':
		num := int(s[index]-'0')*100 + int(s[i1]-'0')*10 + int(s[i2]-'0')
		return 3, s[index] > '0' && 0 <= num && num < 256

	default:
		return -1, false
	}
}

package main

func decodeString(s string) string {
	result, _ := unpack(s, 0, make([]byte, 0, 2000))
	return string(result)
}

func unpack(s string, i int, bytes []byte) ([]byte, int) {
	for i < len(s) && s[i] != ']' {
		if s[i] > '9' {
			bytes, i = append(bytes, s[i]), i+1
			continue
		}

		num, numLength := readInt(s, i)
		prevLength := len(bytes)
		bytes, i = unpack(s, i+numLength+1, bytes)
		bytes = repeat(bytes, prevLength, num)
	}

	return bytes, i + 1
}

func repeat(bytes []byte, from, count int) []byte {
	sub := bytes[from:]
	for i := 1; i < count; i++ {
		bytes = append(bytes, sub...)
	}

	return bytes
}

func readInt(s string, i int) (result, length int) {
	j := i
	for ; s[j] <= '9'; j++ {
		result = result*10 + int(s[j]-'0')
	}

	return result, j - i
}

package kata

func ReverseWords(str string) string {
	result := []byte(str)
	for i, j := 0, 0; i < len(str); i = j {
		for j++; j < len(str) && str[j] != ' '; j++ {
		}

		reverse(result, i, j-1)

		for ; j < len(str) && str[j] == ' '; j++ {
		}
	}

	return string(result)
}

func reverse(b []byte, i, j int) {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i, j = i+1, j-1
	}
}

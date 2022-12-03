package main

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}

	bytes, middleIndex := []byte(palindrome), len(palindrome)>>1
	for i := 0; i < middleIndex; i++ {
		if bytes[i] == 'a' {
			continue
		}

		bytes[i] = 'a'
		return string(bytes)
	}

	bytes[len(bytes)-1] = 'b'
	return string(bytes)
}

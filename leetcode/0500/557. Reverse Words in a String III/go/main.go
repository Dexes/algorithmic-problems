package main

func reverseWords(s string) string {
	var left, right int
	data := []byte(s)

	for left = 0; left < len(s); {
		for ; left < len(s); left++ {
			if data[left] != ' ' {
				break
			}
		}

		for right = left; right < len(s); right++ {
			if data[right] == ' ' {
				break
			}
		}

		reverseSubarray(data, left, right-1)
		left = right + 1
	}

	return string(data)
}

func reverseSubarray(data []byte, left, right int) {
	if right >= len(data) {
		right = len(data) - 1
	}

	for left < right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
}

package main

func reverseStr(s string, k int) string {
	data := []byte(s)
	doubleK := 2 * k

	for i := 0; i < len(data); i += doubleK {
		reverseSubarray(data, i, i+k-1)
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

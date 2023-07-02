package main

var (
	zero  = []byte{'0', '.', '0', '0'}
	bytes = []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '.', '0', '0'}
)

func discountPrices(sentence string, discount int) string {
	result := make([]byte, 0, len(sentence))
	for i, j := 0, 0; i < len(sentence); i++ {
		result = append(result, sentence[i])

		if sentence[i] != '$' || (i > 0 && sentence[i-1] != ' ') {
			continue
		}

		for j = i + 1; j < len(sentence) && '0' <= sentence[j] && sentence[j] <= '9'; j++ {
		}

		if j == i+1 || (j < len(sentence) && sentence[j] != ' ') {
			continue
		}

		result = append(result, applyDiscount(sentence[i+1:j], discount)...)
		i = j - 1
	}

	return string(result)
}

func applyDiscount(price string, discount int) []byte {
	if price == "0" || discount == 100 {
		return zero
	}

	if discount == 0 {
		i := len(bytes) - len(price) - 3
		copy(bytes[i:], price)
		bytes[len(bytes)-2], bytes[len(bytes)-1] = '0', '0'
		return bytes[i:]
	}

	x := 0
	for i := 0; i < len(price); i++ {
		x = (x * 10) + int(price[i]-'0')
	}

	x *= 100 - discount
	i := len(bytes) - 1
	for ; x > 0 || i > 8; i-- {
		if bytes[i] == '.' {
			continue
		}

		bytes[i] = byte(x%10) + '0'
		x /= 10
	}

	return bytes[i+1:]
}

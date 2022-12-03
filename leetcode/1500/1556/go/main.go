package main

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	result, index := make([]byte, 16), 15
	for n > 0 {
		result[index], index = byte(n%10)+'0', index-1
		n /= 10

		if index%4 == 0 && n > 0 {
			result[index], index = '.', index-1
		}
	}

	return string(result[index+1:])
}

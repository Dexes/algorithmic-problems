package main

func intToRoman(num int) string {
	capacity := 15
	result := make([]byte, capacity)

	capacity -= processDigit(num, result[:capacity], 'I', 'V', 'X')
	capacity -= processDigit(num/10, result[:capacity], 'X', 'L', 'C')
	capacity -= processDigit(num/100, result[:capacity], 'C', 'D', 'M')

	for ; num >= 1000; num -= 1000 {
		capacity--
		result[capacity] = 'M'
	}

	return string(result[capacity:])
}

func processDigit(x int, data []byte, one, five, ten byte) int {
	switch x % 10 {
	case 9:
		data[len(data)-2], data[len(data)-1] = one, ten
		return 2
	case 8:
		data[len(data)-4], data[len(data)-3], data[len(data)-2], data[len(data)-1] = five, one, one, one
		return 4
	case 7:
		data[len(data)-3], data[len(data)-2], data[len(data)-1] = five, one, one
		return 3
	case 6:
		data[len(data)-2], data[len(data)-1] = five, one
		return 2
	case 5:
		data[len(data)-1] = five
		return 1
	case 4:
		data[len(data)-2], data[len(data)-1] = one, five
		return 2
	case 3:
		data[len(data)-3], data[len(data)-2], data[len(data)-1] = one, one, one
		return 3
	case 2:
		data[len(data)-2], data[len(data)-1] = one, one
		return 2
	case 1:
		data[len(data)-1] = one
		return 1
	}

	return 0
}

package main

func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		switch {
		case isSingle(data, i):
			i += 1
		case isDouble(data, i):
			i += 2
		case isTriple(data, i):
			i += 3
		case isQuadruple(data, i):
			i += 4
		default:
			return false
		}
	}

	return true
}

func isSingle(data []int, index int) bool {
	return (data[index] & 0b10000000) == 0
}

func isDouble(data []int, index int) bool {
	return index+1 < len(data) &&
		(data[index]&0b11100000) == 0b11000000 &&
		(data[index+1]&0b11000000) == 0b10000000
}

func isTriple(data []int, index int) bool {
	return index+2 < len(data) &&
		(data[index]&0b11110000) == 0b11100000 &&
		(data[index+1]&0b11000000) == 0b10000000 &&
		(data[index+2]&0b11000000) == 0b10000000
}

func isQuadruple(data []int, index int) bool {
	return index+3 < len(data) &&
		(data[index]&0b11111000) == 0b11110000 &&
		(data[index+1]&0b11000000) == 0b10000000 &&
		(data[index+2]&0b11000000) == 0b10000000 &&
		(data[index+3]&0b11000000) == 0b10000000
}

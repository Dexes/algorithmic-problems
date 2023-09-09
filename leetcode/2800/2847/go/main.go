package main

const (
	capacity = 25
	noResult = "-1"
)

func smallestNumber(n int64) string {
	if n < 10 {
		return string(byte(n) + '0')
	}

	result, rIndex := make([]byte, capacity), capacity
	for x := int64(9); x > 1 && n > 1; x-- {
		for ; n%x == 0; n /= x {
			rIndex--
			result[rIndex] = byte(x + '0')
		}
	}

	if n > 10 {
		return noResult
	}

	return string(result[rIndex:])
}

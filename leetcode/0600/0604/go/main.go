package main

type StringIterator struct {
	compressedString string
	currentByte      byte
	currentCount     int
}

func Constructor(compressedString string) StringIterator {
	return StringIterator{compressedString: compressedString}
}

func (x *StringIterator) Next() byte {
	if x.currentCount == 0 {
		if len(x.compressedString) == 0 {
			return ' '
		}

		var index int
		x.currentByte = x.compressedString[0]
		x.currentCount, index = readInt(x.compressedString, 1)
		x.compressedString = x.compressedString[index:]
	}

	x.currentCount--
	return x.currentByte
}

func (x *StringIterator) HasNext() bool {
	return x.currentCount > 0 || len(x.compressedString) > 0
}

func readInt(s string, index int) (int, int) {
	result := 0
	for ; index < len(s) && s[index] <= '9'; index++ {
		result = result*10 + int(s[index]-'0')
	}

	return result, index
}

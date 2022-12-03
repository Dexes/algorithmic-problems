package main

func compress(chars []byte) int {
	rIndex, wIndex, letterNumber := 0, 0, 0
	for ; rIndex < len(chars); rIndex += letterNumber {
		letterNumber = count(chars, rIndex)
		wIndex = write(chars, rIndex, wIndex, letterNumber)
	}

	return wIndex
}

func count(chars []byte, rIndex int) int {
	i := rIndex + 1
	for ; i < len(chars) && chars[i-1] == chars[i]; i++ {
	}

	return i - rIndex
}

func write(chars []byte, rIndex, wIndex, letterNumber int) int {
	chars[wIndex], wIndex = chars[rIndex], wIndex+1
	if letterNumber == 1 {
		return wIndex
	}

	wIndex += countDigits(letterNumber)
	for x := wIndex - 1; letterNumber > 0; letterNumber /= 10 {
		chars[x] = byte(letterNumber%10 + '0')
		x--
	}

	return wIndex
}

func countDigits(x int) int {
	if x < 10 {
		return 1
	}

	if x < 100 {
		return 2
	}

	if x < 1000 {
		return 3
	}

	return 4
}

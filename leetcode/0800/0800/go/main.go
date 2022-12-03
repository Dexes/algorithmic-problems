package main

func similarRGB(color string) string {
	a := round(color[1], color[2])
	b := round(color[3], color[4])
	c := round(color[5], color[6])

	return string([]byte{'#', a, a, b, b, c, c})
}

func round(a, b byte) byte {
	if a == b {
		return a
	}

	a, b = toNum(a), toNum(b)

	var floor, ceil byte
	if a < b {
		floor, ceil = a, a+1
	} else {
		floor, ceil = a-1, a
	}

	num, floorNum, ceilNum := a<<4+b, floor<<4+floor, ceil<<4+ceil
	if num-floorNum < ceilNum-num {
		return toByte(floor)
	}

	return toByte(ceil)
}

func toNum(b byte) byte {
	if b >= 'a' {
		return b - 'a' + 10
	}

	return b - '0'
}

func toByte(b byte) byte {
	if b >= 10 {
		return b - 10 + 'a'
	}

	return b + '0'
}

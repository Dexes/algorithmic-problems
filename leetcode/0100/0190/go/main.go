package main

func reverseBits(num uint32) uint32 {
	for left := 0; left < 16; left++ {
		right := 31 - left
		shift := right - left

		iBit, jBit := num&(1<<left), num&(1<<right)

		if iBit > 0 {
			num |= iBit << shift
		} else {
			num &= ^(1 << right)
		}

		if jBit > 0 {
			num |= jBit >> shift
		} else {
			num &= ^(1 << left)
		}
	}

	return num
}

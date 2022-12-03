package main

func toHex(num int) string {
	bits, i := getBits(num), 0
	for ; bits<<i>>28 == 0 && i < 28; i += 4 {
	}

	result := make([]byte, 0, 8)
	for ; i < 32; i += 4 {
		result = append(result, bitsToHex(bits<<i>>28))
	}

	return string(result)
}

func getBits(n int) uint32 {
	if n < 0 {
		return 0xffffffff ^ uint32(-n-1)
	}

	return uint32(n)
}

func bitsToHex(b uint32) byte {
	if b < 10 {
		return byte('0' + b)
	}

	return byte('a' + b - 10)
}

package main

func hammingWeight(num uint32) int {
	var result uint32
	for ; num > 0; num >>= 1 {
		result += num & 1
	}

	return int(result)
}

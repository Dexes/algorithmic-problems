package main

func rangeBitwiseAnd(left int, right int) int {
	if left == 0 || left == right {
		return left
	}

	result, mask := 0, 1<<findShift(right)
	for mask > 0 && left&mask == right&mask {
		result |= left & mask
		mask >>= 1
	}

	return result
}

func findShift(x int) int {
	for shift := 30; ; shift-- {
		if x&(1<<shift) > 0 {
			return shift
		}
	}
}

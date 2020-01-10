package main

import (
	"bufio"
	"os"
)

func ParseInteger(buf []byte) (int, int) {
	var index, n int
	var v byte

	for index, v = range buf {
		if v < '0' || v > '9' {
			break
		}

		n = n*10 + int(v-'0')
	}

	return n, index
}

func ParseIntegers(buf []byte, length int) ([]int, int) {
	var numSize int
	result := make([]int, length)
	size := 0

	for j := 0; j < length; j++ {
		result[j], numSize = ParseInteger(buf[size:])
		size += numSize + 1
	}

	return result, size
}

func GetIterationsCount(length int) int {
	result := 0
	for i := 0; i < length; i++ {
		result <<= 1
		result |= 1
	}

	return result
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func CalculateHeapsDiff(length int, data []int, bitMask int) int {
	first, second := 0, 0
	pointer := 1

	for i := 0; i < length; i++ {
		if pointer&bitMask == pointer {
			first += data[i]
		} else {
			second += data[i]
		}

		pointer <<= 1
	}

	return Abs(first - second)
}

func GetResult(length int, data []int) int {
	if length == 1 {
		return data[0]
	}

	iterationsCount := GetIterationsCount(length)
	min := 2147483647

	for bitMask := 1; bitMask < iterationsCount; bitMask++ {
		diff := CalculateHeapsDiff(length, data, bitMask)
		if diff < min {
			min = diff
		}
	}

	return min
}

func GetAllInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	input := GetAllInput()
	size, length := ParseInteger(input)
	data, _ := ParseIntegers(input[length+2:], size)

	println(GetResult(size, data))
}

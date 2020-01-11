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

func GetResult(size int, data []int) int {
	candidate := data[0]
	candidateIndex := 0
	candidateCount := 0

	maxCandidateIndex := candidateIndex
	maxCandidateCount := candidateCount

	for index := 1; index < size; index++ {
		candidateCount++
		if data[index] < candidate {
			if candidateCount > maxCandidateCount {
				maxCandidateCount = candidateCount
				maxCandidateIndex = candidateIndex
			}

			candidate = data[index]
			candidateCount = 1
			candidateIndex = index
		}
	}

	if candidateCount > maxCandidateCount {
		maxCandidateIndex = candidateIndex
	}

	return maxCandidateIndex + 1
}

func GetInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	input := GetInput()
	size, l := ParseInteger(input)
	data, _ := ParseIntegers(input[l+2:], size)

	println(GetResult(size, data))
}

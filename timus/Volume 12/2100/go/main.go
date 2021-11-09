package main

import (
	"bufio"
	"os"
)

func ParseInteger(buf []byte) (n, index int) {
	var v byte
	for index, v = range buf {
		if v < '0' || v > '9' {
			break
		}

		n = n*10 + int(v-'0')
	}

	return
}

func GetFirstRowSize(input []byte) int {
	length := len(input)
	result := 0

	for ; result < length; result++ {
		if input[result] == '\n' {
			break
		}
	}

	return result + 1
}

func ParseInput(input []byte) [][]byte {
	count, pointer := ParseInteger(input)
	result := make([][]byte, count)

	pointer += 2
	for i := 0; i < count; i++ {
		size := GetFirstRowSize(input[pointer:])
		result[i] = input[pointer : pointer+size-2]
		pointer += size
	}

	return result
}

func GetResult(data [][]byte) int {
	result := 2
	for _, buf := range data {
		length := len(buf)
		if length > 4 && buf[length-4] == '+' {
			result += 2
		} else {
			result++
		}
	}

	if result == 13 {
		return 1400
	}

	return result * 100
}

func GetInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	println(GetResult(ParseInput(GetInput())))
}

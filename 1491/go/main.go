package main

import (
	"bufio"
	"os"
	"strconv"
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

func MakeFenwick(size int, data [][]int) []int {
	fenwick := make([]int, size)
	for _, row := range data {
		left := row[0] - 1
		right := row[1]

		fenwick[left] += row[2]
		if right < size {
			fenwick[right] -= row[2]
		}
	}

	return fenwick
}

func GetResult(size int, data [][]int) []int {
	fenwick := MakeFenwick(size, data)
	for i := 1; i < size; i++ {
		fenwick[i] += fenwick[i-1]
	}

	return fenwick
}

func GetAllInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	input := GetAllInput()
	size, length := ParseInteger(input)
	data := make([][]int, size+1)

	pointer := length + 2
	for i := 0; i <= size; i++ {
		data[i], length = ParseIntegers(input[pointer:], 3)
		pointer += length + 1
	}

	result := GetResult(size, data)

	print(strconv.Itoa(result[0]), " ")
	for i := 1; i < size; i++ {
		print(result[i], " ")
	}
	println()
}

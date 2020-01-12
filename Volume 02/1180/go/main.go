package main

import (
	"bufio"
	"os"
)

func GetResult(data []byte) (int, int) {
	var sum int
	for _, b := range data {
		if b < '0' || b > '9' {
			break
		}

		sum += int(b)
	}

	sum %= 3
	if sum != 0 {
		return 1, sum
	}

	return 2, 0
}

func GetInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	first, second := GetResult(GetInput())
	if first == 1 {
		print("1\n", second)
	} else {
		print(first)
	}
}

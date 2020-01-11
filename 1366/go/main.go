package main

import (
	"bufio"
	"math/big"
	"os"
)

func ToInt(data []byte) int64 {
	var n int64
	for _, b := range data {
		if b < '0' || b > '9' {
			break
		}

		n = n*10 + int64(b-'0')
	}

	return n
}

func GetResult(amount int64) string {
	// http://mathworld.wolfram.com/Subfactorial.html

	current := big.NewInt(0)
	increase := big.NewInt(1)
	decrease := big.NewInt(-1)
	even := false

	for i := int64(0); i <= amount; i++ {
		current.Mul(current, big.NewInt(i))
		if even = !even; even {
			current.Add(current, increase)
		} else {
			current.Add(current, decrease)
		}
	}

	return current.String()
}

func GetInput() []byte {
	in := bufio.NewReader(os.Stdin)
	bytes, _ := in.ReadBytes(0)

	return bytes
}

func main() {
	println(GetResult(ToInt(GetInput())))
}

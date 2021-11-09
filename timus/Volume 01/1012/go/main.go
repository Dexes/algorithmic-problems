package main

import (
	"fmt"
	"math/big"
)

func GetInput() (int, int64) {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	return n, int64(k)
}

func main() {
	var n, intK = GetInput()

	var k = big.NewInt(intK - 1)
	var first = big.NewInt(intK - 1)
	var second = big.NewInt(0).Mul(first, big.NewInt(intK))

	for i := 2; i < n; i++ {
		var tmp = big.NewInt(0).Add(first, second)
		tmp = tmp.Mul(tmp, k)

		first = second
		second = tmp
	}

	println(second.String())
}

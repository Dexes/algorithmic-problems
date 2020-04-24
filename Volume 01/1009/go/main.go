package main

import "fmt"

func GetInput() (int, int) {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	return n, k
}

func main() {
	var n, k = GetInput()
	var first = k - 1
	var second = first * k

	for i := 2; i < n; i++ {
		first, second = second, (first+second)*(k-1)
	}

	println(second)
}

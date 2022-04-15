package main

import (
	"fmt"
)

var steps = [][]byte{{1, 2}, {1, 254}, {2, 1}, {2, 255}, {254, 1}, {254, 255}, {255, 2}, {255, 254}}

func count(a, b byte) (result int) {
	for _, step := range steps {
		if a+step[0] < 8 && b+step[1] < 8 {
			result++
		}
	}

	return result
}

func main() {
	var (
		n int
		s string
	)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		fmt.Println(count(s[0]-'a', s[1]-'1'))
	}
}

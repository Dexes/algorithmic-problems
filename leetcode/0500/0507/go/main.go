package main

import (
	"math"
)

func checkPerfectNumber(num int) bool {
	sum, sqrt := 0, int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}

	return num > 1 && sum+1 == num
}

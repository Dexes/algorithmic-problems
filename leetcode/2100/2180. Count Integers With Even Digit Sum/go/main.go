package main

func countEven(num int) int {
	sum, x := 0, num
	for x > 0 {
		sum += x % 10
		x /= 10
	}

	if sum&1 == 1 {
		return (num - 1) / 2
	}

	return num / 2
}

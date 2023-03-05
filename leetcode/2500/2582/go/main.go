package main

func passThePillow(n int, time int) int {
	rounds, reminder := time/(n-1), time%(n-1)
	if (rounds & 1) == 1 {
		return n - reminder
	}

	return reminder + 1
}

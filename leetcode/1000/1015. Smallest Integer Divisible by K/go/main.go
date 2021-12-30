package main

func smallestRepunitDivByK(k int) int {
	for num, count := 0, 1; count <= k; count++ {
		num = (num*10 + 1) % k
		if k == 0 {
			return count
		}
	}

	return -1
}

package main

func minPartitions(n string) int {
	result := byte(0)
	for i := 0; i < len(n); i++ {
		if n[i] > result {
			result = n[i]
		}
	}

	return int(result - '0')
}

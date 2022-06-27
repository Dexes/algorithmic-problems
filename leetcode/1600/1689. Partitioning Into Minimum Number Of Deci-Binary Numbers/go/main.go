package main

func minPartitions(n string) int {
	result := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > result {
			result = n[i]
		}
	}

	return int(result - '0')
}

package main

func numJewelsInStones(jewels string, stones string) int {
	gems := make(map[byte]bool)
	for i := 0; i < len(jewels); i++ {
		gems[jewels[i]] = true
	}

	result := 0
	for i := 0; i < len(stones); i++ {
		if gems[stones[i]] {
			result++
		}
	}

	return result
}

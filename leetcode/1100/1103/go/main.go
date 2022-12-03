package main

func distributeCandies(candies int, numPeople int) []int {
	result, x, i := make([]int, numPeople), 1, 0
	for {
		for i = 0; i < len(result) && x < candies; i++ {
			result[i], candies, x = result[i]+x, candies-x, x+1
		}

		if x >= candies {
			if i == len(result) {
				i = 0
			}

			result[i] += candies
			return result
		}
	}
}

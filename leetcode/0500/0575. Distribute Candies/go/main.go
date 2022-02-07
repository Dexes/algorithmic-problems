package main

func distributeCandies(candyType []int) int {
	types := make(map[int]struct{})
	for i := 0; i < len(candyType); i++ {
		types[candyType[i]] = struct{}{}
	}

	if len(types) > len(candyType)/2 {
		return len(candyType) / 2
	}

	return len(types)
}

package main

func bestHand(ranks []int, suits []byte) string {
	if suits[0] == suits[1] && suits[0] == suits[2] && suits[0] == suits[3] && suits[0] == suits[4] {
		return "Flush"
	}

	counter, pair := [14]int{}, false
	for i := 0; i < len(ranks); i++ {
		counter[ranks[i]]++

		if counter[ranks[i]] == 3 {
			return "Three of a Kind"
		}

		if counter[ranks[i]] == 2 {
			pair = true
		}
	}

	if pair {
		return "Pair"
	}

	return "High Card"
}

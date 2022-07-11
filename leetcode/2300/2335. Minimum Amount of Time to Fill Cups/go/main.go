package main

func fillCups(amount []int) (result int) {
	if amount[1] > amount[0] && amount[1] > amount[2] {
		amount[0], amount[1] = amount[1], amount[0]
	}

	if amount[2] > amount[0] && amount[2] > amount[1] {
		amount[0], amount[2] = amount[2], amount[0]
	}

	if amount[0] >= amount[1]+amount[2] {
		return amount[0]
	}

	return (amount[0] + amount[1] + amount[2] + 1) / 2
}

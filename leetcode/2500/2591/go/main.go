package main

func distMoney(money int, children int) int {
	if children > money {
		return -1
	}

	money -= children
	result, reminder := money/7, money%7

	if result > children {
		return children - 1
	}

	if reminder > 0 && result == children {
		return result - 1
	}

	if reminder == 3 && result+1 == children {
		return result - 1
	}

	return result
}

package main

func distanceTraveled(mainTank int, additionalTank int) int {
	return (mainTank + min((mainTank-1)/4, additionalTank)) * 10
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

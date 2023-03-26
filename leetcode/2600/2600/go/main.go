package main

func kItemsWithMaximumSum(numOnes, numZeros, _, k int) int {
	if k <= numOnes {
		return k
	}

	if k -= numOnes; k <= numZeros {
		return numOnes
	}

	return numOnes - (k - numZeros)
}

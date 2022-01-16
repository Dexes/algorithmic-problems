package main

var factorials = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func getPermutation(n int, k int) string {
	result, usedNumbers := make([]byte, 0, n), make([]bool, n)
	for i := n - 1; i >= 0; i-- {
		indexNumber, remainder := k/factorials[i], k%factorials[i]
		if remainder == 0 {
			result = append(result, findNumber(usedNumbers, indexNumber-1)+'0')
			k = factorials[i]
		} else {
			result = append(result, findNumber(usedNumbers, indexNumber)+'0')
			k = remainder
		}
	}

	return string(result)
}

func findNumber(usedNumbers []bool, indexNumber int) byte {
	result := byte(0)
	for counter := -1; counter < indexNumber; result++ {
		if !usedNumbers[result] {
			counter++
		}
	}

	usedNumbers[result-1] = true
	return result
}

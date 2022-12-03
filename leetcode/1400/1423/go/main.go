package main

func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	result, left, right := sum, k-1, len(cardPoints)-1
	for i := 0; i < k; i++ {
		if sum += cardPoints[right-i] - cardPoints[left-i]; sum > result {
			result = sum
		}
	}

	return result
}

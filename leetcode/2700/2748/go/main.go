package main

func countBeautifulPairs(nums []int) (result int) {
	for i, n := range nums {
		for ; n > 9; n /= 10 {
		}

		for i++; i < len(nums); i++ {
			if coprime(nums[i]%10, n) {
				result++
			}
		}
	}

	return result
}

func coprime(a, b int) bool {
	if (a%2 == 0 && b%2 == 0) ||
		(a%3 == 0 && b%3 == 0) ||
		(a == 5 && b == 5) ||
		(a == 7 && b == 7) {
		return false
	}

	return true
}

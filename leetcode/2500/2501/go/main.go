package main

func longestSquareStreak(nums []int) (result int) {
	lengths := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		lengths[nums[i]] = 1
	}

	for i := 0; i < len(nums); i++ {
		if x := streakLength(nums[i], lengths); x > result {
			result = x
		}
	}

	if result == 1 {
		return -1
	}

	return result
}

func streakLength(num int, lengths map[int]int) (result int) {
	for x := num; lengths[x] > 0; x *= x {
		if lengths[x] > 1 {
			lengths[num] = result + lengths[x]
			return lengths[num]
		}

		lengths[x] = 0
		result++
	}

	lengths[num] = result
	return result
}

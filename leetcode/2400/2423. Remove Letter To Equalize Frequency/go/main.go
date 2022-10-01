package main

func equalFrequency(word string) bool {
	frequencies, maxFrequency := make([]int, 123), 0
	for i := 0; i < len(word); i++ {
		if frequencies[word[i]]++; frequencies[word[i]] > maxFrequency {
			maxFrequency = frequencies[word[i]]
		}
	}

	frequencies = compress(frequencies)

	return tryRemoveSingleLetter(frequencies, maxFrequency) || tryRemoveMaxLetter(frequencies, maxFrequency)
}

func compress(nums []int) []int {
	wIndex, rIndex := 'a', 'a'
	for ; rIndex <= 'z'; rIndex++ {
		if nums[rIndex] > 0 {
			nums[wIndex] = nums[rIndex]
			wIndex++
		}
	}

	return nums['a':wIndex]
}

func tryRemoveSingleLetter(frequencies []int, maxFrequency int) bool {
	if maxFrequency == 1 {
		return true
	}

	removed := 0
	for i := 0; i < len(frequencies); i++ {
		if frequencies[i] == maxFrequency {
			continue
		}

		if frequencies[i] == 1 {
			if removed++; removed > 1 {
				return false
			}

			continue
		}

		return false
	}

	return removed == 1
}

func tryRemoveMaxLetter(frequencies []int, maxFrequency int) bool {
	maxCount := 0
	for i := 0; i < len(frequencies); i++ {
		if frequencies[i] == maxFrequency {
			if maxCount++; maxCount > 1 {
				return false
			}

			continue
		}

		if frequencies[i]+1 != maxFrequency {
			return false
		}
	}

	return true
}

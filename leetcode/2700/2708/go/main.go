package main

func maxStrength(nums []int) int64 {
	negative, negativeCount, negativeMax := 1, 0, 0
	positive, positiveCount := 1, 0
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		}

		if nums[i] > 0 {
			positive *= nums[i]
			positiveCount++
			continue
		}

		negative *= -nums[i]
		negativeCount++
		if negativeMax == 0 || nums[i] > negativeMax {
			negativeMax = nums[i]
		}
	}

	if positiveCount > 0 {
		if (negativeCount % 2) == 1 {
			negative /= -negativeMax
		}

		return int64(positive * negative)
	}

	if negativeCount < 2 {
		if zeroCount > 0 {
			return 0
		}

		return int64(negativeMax)
	}

	if (negativeCount % 2) == 1 {
		negative /= -negativeMax
	}

	return int64(negative)
}

package main

func sortedSquares(nums []int) []int {
	positiveIndex := firstPositiveIndex(nums)
	negativeIndex := positiveIndex - 1
	square(nums)

	result := make([]int, 0, len(nums))
	for negativeIndex >= 0 || positiveIndex < len(nums) {
		if negativeIndex < 0 {
			result = append(result, nums[positiveIndex])
			positiveIndex++
			continue
		}

		if positiveIndex == len(nums) || nums[negativeIndex] < nums[positiveIndex] {
			result = append(result, nums[negativeIndex])
			negativeIndex--
			continue
		}

		result = append(result, nums[positiveIndex])
		positiveIndex++
	}

	return result
}

func square(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
}

func firstPositiveIndex(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		middle := left + (right-left)/2

		if nums[middle] >= 0 {
			right = middle
		} else {
			left = middle + 1
		}
	}

	if nums[left] < 0 {
		return left + 1
	}

	return left
}

package main

func countBalls(lowLimit int, highLimit int) int {
	boxes := make([]int, 46)
	for ; lowLimit <= highLimit; lowLimit++ {
		boxes[sum(lowLimit)]++
	}

	return max(boxes)
}

func sum(n int) int {
	result := 0
	for n > 0 {
		result += n % 10
		n /= 10
	}

	return result
}

func max(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}

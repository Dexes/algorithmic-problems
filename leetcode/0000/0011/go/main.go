package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	for left < right {
		if height[left] < height[right] {
			result = max(result, height[left]*(right-left))
			left++
		} else {
			result = max(result, height[right]*(right-left))
			right--
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

package main

func getMinDistance(nums []int, target int, index int) int {
	left, right := index, index
	for ; left >= 0 && nums[left] != target; left-- {
	}

	for ; right < len(nums) && nums[right] != target; right++ {
	}

	switch lDistance, rDistance := index-left, right-index; {
	case left < 0:
		return rDistance
	case right == len(nums):
		return lDistance
	case rDistance < lDistance:
		return rDistance
	default:
		return lDistance
	}
}

package main

type HitCounter []int

func Constructor() HitCounter {
	return make([]int, 0, 300)
}

func (x *HitCounter) Hit(timestamp int) {
	*x = append(*x, timestamp)
}

func (x *HitCounter) GetHits(timestamp int) int {
	nums, t := *x, timestamp-299
	if len(nums) == 0 {
		return 0
	}

	if left := search(nums, t); nums[left] >= t {
		return len(nums) - left
	}

	return 0
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if middle := left + (right-left)/2; nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return left
}

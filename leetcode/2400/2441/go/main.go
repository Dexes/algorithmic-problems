package main

func findMaxK(nums []int) int {
	max, set := -1, make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := set[-nums[i]]; !ok {
			set[nums[i]] = struct{}{}
			continue
		}

		if x := abs(nums[i]); x > max {
			max = x
		}
	}

	return max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

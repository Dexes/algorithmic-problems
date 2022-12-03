package main

func arraySign(nums []int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}

		if nums[i] < 0 {
			counter++
		}
	}

	if counter&1 == 1 {
		return -1
	}

	return 1
}

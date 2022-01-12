package main

func majorityElement(nums []int) int {
	balance, element := 0, 0
	for i := 0; i < len(nums); i++ {
		if balance == 0 {
			balance, element = 1, nums[i]
			continue
		}

		if nums[i] == element {
			balance++
		} else {
			balance--
		}
	}

	return element
}

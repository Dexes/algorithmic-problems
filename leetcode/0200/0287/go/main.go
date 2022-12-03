package main

func findDuplicate(nums []int) (fast int) {
	for slow := 0; ; {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	for slow := 0; slow != fast; {
		slow, fast = nums[slow], nums[fast]
	}

	return fast
}

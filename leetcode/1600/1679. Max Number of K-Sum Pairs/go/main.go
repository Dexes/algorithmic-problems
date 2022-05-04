package main

func maxOperations(nums []int, k int) (result int) {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if x := k - nums[i]; counter[x] > 0 {
			counter[x]--
			result++
			continue
		}

		counter[nums[i]]++
	}

	return result
}

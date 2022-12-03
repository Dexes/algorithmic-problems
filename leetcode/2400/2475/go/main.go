package main

func unequalTriplets(nums []int) (result int) {
	pairs, counter := 0, make([]int, 1001)
	for i := 0; i < len(nums); i++ {
		left := i - counter[nums[i]]
		result += pairs - counter[nums[i]]*left
		pairs += left

		counter[nums[i]]++
	}

	return result
}

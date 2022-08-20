package main

func arithmeticTriplets(nums []int, diff int) (result int) {
	double := diff << 1
	set := make([]bool, 201)
	set[nums[0]], set[nums[1]] = true, true

	for i := 2; i < len(nums); i++ {
		first, second := nums[i]-diff, nums[i]-double
		if second >= 0 && set[first] && set[second] {
			result++
		}

		set[nums[i]] = true
	}

	return result
}

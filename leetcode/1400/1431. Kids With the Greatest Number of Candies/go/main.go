package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := maximum(candies)

	for i := 0; i < len(candies); i++ {
		result[i] = candies[i]+extraCandies >= max
	}

	return result
}

func maximum(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}

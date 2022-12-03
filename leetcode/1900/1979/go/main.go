package main

func findGCD(nums []int) int {
	return GCD(getMinMax(nums))
}

func getMinMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return min, max
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

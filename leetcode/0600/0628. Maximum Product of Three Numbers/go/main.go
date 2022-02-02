package main

func maximumProduct(nums []int) int {
	a, b := 1000, 1000
	x, y, z := -1000, -1000, -1000
	for i := 0; i < len(nums); i++ {
		if nums[i] < a {
			a, b = nums[i], a
		} else if nums[i] < b {
			b = nums[i]
		}

		if nums[i] > z {
			z, y, x = nums[i], z, y
		} else if nums[i] > y {
			y, x = nums[i], y
		} else if nums[i] > x {
			x = nums[i]
		}
	}

	first := a * b * z
	second := z * y * x
	if first > second {
		return first
	}

	return second
}

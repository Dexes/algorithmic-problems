package main

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		image[i] = reverse(image[i])
	}

	return image
}

func reverse(nums []int) []int {
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		if nums[left] != nums[right] {
			continue
		}

		if nums[left] == 0 {
			nums[left], nums[right] = 1, 1
		} else {
			nums[left], nums[right] = 0, 0
		}
	}

	if len(nums)&1 == 1 {
		if nums[len(nums)>>1] == 1 {
			nums[len(nums)>>1] = 0
		} else {
			nums[len(nums)>>1] = 1
		}
	}

	return nums
}

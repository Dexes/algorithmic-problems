package main

func pivotArray(nums []int, pivot int) []int {
	tail := make([]int, 0, len(nums)/4)
	rIndex, wIndex, counter := 0, 0, 0

	for ; rIndex < len(nums); rIndex++ {
		switch {
		case nums[rIndex] < pivot:
			nums[wIndex] = nums[rIndex]
			wIndex++
		case nums[rIndex] == pivot:
			counter++
		default:
			tail = append(tail, nums[rIndex])
		}
	}

	for rIndex = 0; rIndex < counter; rIndex++ {
		nums[wIndex] = pivot
		wIndex++
	}

	for rIndex = 0; rIndex < len(tail); rIndex++ {
		nums[wIndex] = tail[rIndex]
		wIndex++
	}

	return nums
}

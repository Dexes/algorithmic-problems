package main

func sortEvenOdd(nums []int) []int {
	evenRight, oddRight := len(nums)-2, len(nums)-1
	if evenRight&1 == 1 {
		evenRight, oddRight = oddRight, evenRight
	}

	quickSort(nums, 0, evenRight, 2, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	quickSort(nums, 1, oddRight, 2, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums
}

func quickSort(arr []int, low, high, step int, comparer func(i, j int) bool) {
	if low < high {
		p := partition(arr, low, high, step, comparer)
		quickSort(arr, low, p-step, step, comparer)
		quickSort(arr, p+step, high, step, comparer)
	}
}

func partition(arr []int, low, high, step int, comparer func(i, j int) bool) int {
	for j := low; j < high; j += step {
		if comparer(j, high) {
			arr[low], arr[j] = arr[j], arr[low]
			low += step
		}
	}

	arr[low], arr[high] = arr[high], arr[low]
	return low
}

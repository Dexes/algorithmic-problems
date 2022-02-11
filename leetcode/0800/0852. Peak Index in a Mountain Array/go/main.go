package main

func peakIndexInMountainArray(arr []int) int {
	i := 1
	for ; i < len(arr) && arr[i-1] < arr[i]; i++ {
	}

	return i - 1
}

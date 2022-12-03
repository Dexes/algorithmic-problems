package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 1
	for ; i < len(arr) && arr[i-1] < arr[i]; i++ {
	}

	if i == 1 || i == len(arr) {
		return false
	}

	for ; i < len(arr) && arr[i-1] > arr[i]; i++ {
	}

	return i == len(arr)
}

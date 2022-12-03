package main

func findKthPositive(arr []int, k int) int {
	if arr[0] > k {
		return k
	}

	if arr[len(arr)-1]-len(arr) < k {
		return k + len(arr)
	}

	for left, right := 1, len(arr)-2; ; {
		switch middle := left + (right-left)/2; {
		case k > arr[middle]-(middle+1):
			left = middle + 1
		case k <= arr[middle-1]-middle:
			right = middle - 1
		default:
			return k + middle
		}
	}
}

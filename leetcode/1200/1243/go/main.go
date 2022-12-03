package main

func transformArray(arr []int) []int {
	for rightIndex, changed := len(arr)-1, true; changed; {
		changed = false
		for i, prev := 1, arr[0]; i < rightIndex; i++ {
			current, next := arr[i], arr[i+1]
			if prev > current && current < next {
				arr[i]++
				changed = true
			} else if prev < current && current > next {
				arr[i]--
				changed = true
			}

			prev = current
		}
	}

	return arr
}

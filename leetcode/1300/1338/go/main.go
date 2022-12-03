package main

func minSetSize(arr []int) (result int) {
	frequencies := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		frequencies[arr[i]]++
		arr[i] = 0
	}

	for _, frequency := range frequencies {
		arr[frequency-1]++
	}

	for needRemove, i := len(arr)/2, len(arr)-1; ; i-- {
		if arr[i] == 0 {
			continue
		}

		frequency, count := i+1, arr[i]
		if x := frequency * count; x < needRemove {
			needRemove -= x
			result += count
			continue
		}

		return result + search(frequency, count, needRemove)
	}
}

func search(frequency, count, needRemove int) int {
	left, right := 0, count
	for left < right {
		if middle := left + (right-left)/2; middle*frequency < needRemove {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return left
}

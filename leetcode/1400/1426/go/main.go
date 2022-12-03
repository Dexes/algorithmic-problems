package main

func countElements(arr []int) (result int) {
	counter := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		counter[arr[i]]++
	}

	for num, count := range counter {
		if _, ok := counter[num+1]; ok {
			result += count
		}
	}

	return result
}

package main

func kthDistinct(arr []string, k int) string {
	letters := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		letters[arr[i]]++
	}

	number := 0
	for i := 0; i < len(arr); i++ {
		if letters[arr[i]] == 1 {
			number++
			if number == k {
				return arr[i]
			}
		}
	}

	return ""
}

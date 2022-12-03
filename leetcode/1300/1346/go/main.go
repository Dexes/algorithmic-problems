package main

func checkIfExist(arr []int) bool {
	index := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if index[arr[i]*2] || arr[i]&1 == 0 && index[arr[i]>>1] {
			return true
		}

		index[arr[i]] = true
	}

	return false
}

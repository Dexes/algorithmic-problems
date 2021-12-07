package main

func canBeEqual(target []int, arr []int) bool {
	balance := make(map[int]int)
	for i := 0; i < len(target); i++ {
		balance[target[i]]++
		balance[arr[i]]--
	}

	for _, value := range balance {
		if value != 0 {
			return false
		}
	}

	return true
}

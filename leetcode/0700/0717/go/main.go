package main

func isOneBitCharacter(bits []int) bool {
	result := true
	for i := len(bits) - 2; i >= 0 && bits[i] == 1; i-- {
		result = !result
	}

	return result
}

package main

func buildArray(target []int, n int) []string {
	result, current := make([]string, 0, 150), 1
	for i := 0; i < len(target); i++ {
		for ; current < target[i]; current++ {
			result = append(result, "Push", "Pop")
		}

		result = append(result, "Push")
		current++
	}

	return result
}

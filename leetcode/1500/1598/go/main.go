package main

func minOperations(logs []string) int {
	result := 0

	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" {
			if result > 0 {
				result--
			}
		} else if logs[i] != "./" {
			result++
		}
	}

	return result
}

package main

func generatePossibleNextMoves(currentState string) []string {
	result, bytes := make([]string, 0, len(currentState)/2), []byte(currentState)
	for i := 1; i < len(currentState); i++ {
		if currentState[i] == '-' || currentState[i-1] != currentState[i] {
			continue
		}

		bytes[i-1], bytes[i] = '-', '-'
		result = append(result, string(bytes))
		bytes[i-1], bytes[i] = '+', '+'
	}

	return result
}

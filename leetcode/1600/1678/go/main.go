package main

func interpret(command string) string {
	result := make([]byte, 0, len(command))
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			result = append(result, 'G')
			continue
		}

		if command[i+1] == ')' {
			result = append(result, 'o')
			i += 1
			continue
		}

		result = append(result, 'a', 'l')
		i += 3
	}

	return string(result)
}

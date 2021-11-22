package main

func makePermutations(permutations []string, current []byte, index, balance int) []string {
	if balance > len(current)-index {
		return permutations
	}

	if index == len(current) {
		return append(permutations, string(current))
	}

	current[index] = '('
	permutations = makePermutations(permutations, current, index+1, balance+1)

	if balance > 0 {
		current[index] = ')'
		permutations = makePermutations(permutations, current, index+1, balance-1)
	}

	return permutations
}

func generateParenthesis(n int) []string {
	return makePermutations(
		make([]string, 0, 1430),
		make([]byte, n*2),
		0, 0,
	)
}

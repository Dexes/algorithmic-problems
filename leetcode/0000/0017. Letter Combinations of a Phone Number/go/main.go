package main

var letters = map[uint8][]uint8{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func makePermutations(digits string, index int, data []uint8, result []string) []string {
	if index == len(data) {
		return append(result, string(data))
	}

	for _, letter := range letters[digits[index]] {
		data[index] = letter
		result = makePermutations(digits, index+1, data, result)
	}

	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := make([]string, 0)
	data := make([]uint8, len(digits))

	result = makePermutations(digits, 0, data, result)

	return result
}

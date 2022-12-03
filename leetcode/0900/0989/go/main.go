package main

func addToArrayForm(num []int, k int) []int {
	return sum(num, toArray(k))
}

func sum(first []int, second []int) []int {
	if len(second) > len(first) {
		first, second = second, first
	}

	firstIndex, secondIndex := len(first)-1, len(second)-1
	digit, buffer := 0, 0

	for firstIndex >= 0 && (secondIndex >= 0 || buffer > 0) {
		if secondIndex < 0 {
			digit = 0
		} else {
			digit = second[secondIndex]
		}

		first[firstIndex] += buffer + digit
		if first[firstIndex] >= 10 {
			first[firstIndex] -= 10
			buffer = 1
		} else {
			buffer = 0
		}

		firstIndex--
		secondIndex--
	}

	if buffer == 0 {
		return first
	}

	return append([]int{1}, first...)
}

func toArray(num int) []int {
	result := make([]int, 10)
	i := len(result) - 1
	for ; num > 0; i-- {
		result[i] = num % 10
		num /= 10
	}

	return result[i+1:]
}

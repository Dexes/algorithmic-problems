package main

type CategoryHandler interface {
	HaveSameCategory(int, int) bool
}

func numberOfCategories(n int, handler CategoryHandler) int {
	categories := 1
	elements := make([]int, 1, n)

	for element := 1; element < n; element++ {
		if haveSameCategory(elements, element, handler) {
			continue
		}

		elements = append(elements, element)
		categories++
	}

	return categories
}

func haveSameCategory(elements []int, element int, handler CategoryHandler) bool {
	for i := 0; i < len(elements); i++ {
		if handler.HaveSameCategory(elements[i], element) {
			return true
		}
	}

	return false
}

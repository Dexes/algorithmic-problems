package main

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	result := make([][]string, 0, len(searchWord))
	for i := 0; i < len(searchWord); i++ {
		products = filter(products, i, searchWord[i])
		result = append(result, products[:min(len(products), 3)])
	}

	return result
}

func filter(products []string, position int, b byte) []string {
	i, j := 0, 0
	for ; i < len(products) && (position == len(products[i]) || products[i][position] != b); i++ {
	}

	for j = i; j < len(products) && products[j][position] == b; j++ {
	}

	return products[i:j]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

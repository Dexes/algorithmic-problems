package main

func isValidSerialization(preorder string) bool {
	balance := 1
	for i := 0; i < len(preorder); i = skipValue(preorder, i) {
		if balance--; balance < 0 {
			return false
		}

		if preorder[i] >= '0' {
			balance += 2
		}
	}

	return balance == 0
}

func skipValue(s string, i int) int {
	for ; i < len(s) && s[i] != ','; i++ {
	}

	return i + 1
}

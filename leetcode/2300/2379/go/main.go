package main

func minimumRecolors(blocks string, k int) (result int) {
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			result++
		}
	}

	current := result
	for i, j := 0, k; j < len(blocks); i, j = i+1, j+1 {
		if blocks[i] == blocks[j] {
			continue
		}

		if blocks[i] == 'W' {
			current--
			if current < result {
				result = current
			}
		} else {
			current++
		}
	}

	return result
}

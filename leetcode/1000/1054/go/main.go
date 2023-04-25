package main

func rearrangeBarcodes(barcodes []int) []int {
	if len(barcodes) < 3 {
		return barcodes
	}

	counter, num, cnt := count(barcodes)
	index := set(barcodes, 0, num, cnt)
	for num, cnt = range counter {
		index = set(barcodes, index, num, cnt)
	}

	return barcodes
}

func count(barcodes []int) (map[int]int, int, int) {
	result, num, max, current := make(map[int]int), 0, 0, 0
	for i := 0; i < len(barcodes); i++ {
		if current = result[barcodes[i]] + 1; current > max {
			num, max = barcodes[i], current
		}

		result[barcodes[i]] = current
	}

	delete(result, num)
	return result, num, max
}

func set(barcodes []int, index, num, count int) int {
	for ; count > 0; count-- {
		if index >= len(barcodes) {
			index = 1
		}

		barcodes[index] = num
		index += 2
	}

	return index
}

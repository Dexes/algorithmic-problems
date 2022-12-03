package main

func cellsInRange(s string) []string {
	l1, l2, d1, d2 := s[0], s[3], s[1], s[4]
	result, cell := make([]string, 0, (l2-l1+1)*(d2-d1+1)), []byte{0, 0}
	for ; l1 <= l2; l1++ {
		for i := d1; i <= d2; i++ {
			cell[0], cell[1] = l1, i
			result = append(result, string(cell))
		}
	}

	return result
}

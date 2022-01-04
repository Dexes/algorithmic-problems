package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]byte, 0, len(s))
	diagonalDst, verticalDst := (numRows-1)<<1, 0
	for row := 0; row < numRows; row++ {
		evenIteration := true
		for i := row; i < len(s); {
			result = append(result, s[i])

			if verticalDst == 0 || evenIteration && diagonalDst != 0 {
				i += diagonalDst
			} else {
				i += verticalDst
			}

			evenIteration = !evenIteration
		}

		diagonalDst -= 2
		verticalDst += 2
	}

	return string(result)
}

package main

func canThreePartsEqualSum(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}

	preLastIndex, lastIndex := len(arr)-2, len(arr)-1
	if arr[lastIndex]%3 > 0 {
		return false
	}

	oneThird, twoThirds := arr[lastIndex]/3, arr[lastIndex]/3*2
	for i := 0; i < preLastIndex; i++ {
		if arr[i] != oneThird {
			continue
		}

		for j := i + 1; j < lastIndex; j++ {
			if arr[j] == twoThirds && arr[lastIndex]-arr[j] == oneThird {
				return true
			}
		}
	}

	return false
}

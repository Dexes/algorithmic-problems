package main

func totalFruit(fruits []int) int {
	i := 1
	for ; i < len(fruits) && fruits[0] == fruits[i]; i++ {
	}

	if i == len(fruits) {
		return i
	}

	aIndex, aCount, bIndex, bCount := i-1, i, i, 1
	aBefore, bBefore := aCount, 0
	result := 0
	for i++; i < len(fruits); i++ {
		if fruits[i] == fruits[aIndex] {
			aIndex, aCount, bBefore = i, aCount+1, bCount
			continue
		}

		if fruits[i] == fruits[bIndex] {
			bIndex, bCount, aBefore = i, bCount+1, aCount
			continue
		}

		if x := aCount + bCount; x > result {
			result = x
		}

		if aIndex < bIndex {
			aIndex, aCount, bCount = i, 1, bCount-bBefore
			aBefore, bBefore = 0, bCount
		} else {
			bIndex, bCount, aCount = i, 1, aCount-aBefore
			aBefore, bBefore = aCount, 0
		}
	}

	if x := aCount + bCount; x > result {
		return x
	}

	return result
}

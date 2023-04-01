package main

func minNumber(first, second []int) int {
	fDigits, sDigits := toDigits(first), toDigits(second)
	for i := 1; i < 10; i++ {
		if fDigits[i] && sDigits[i] {
			return i
		}
	}

	fMin, sMin := 1, 1
	for ; !fDigits[fMin]; fMin++ {
	}

	for ; !sDigits[sMin]; sMin++ {
	}

	if sMin < fMin {
		fMin, sMin = sMin, fMin
	}

	return fMin*10 + sMin
}

func toDigits(nums []int) (result [10]bool) {
	for i := 0; i < len(nums); i++ {
		result[nums[i]] = true
	}

	return result
}

package main

func garbageCollection(garbage []string, travel []int) int {
	volume := len(garbage[0])
	mIndex, pIndex, gIndex := 0, 0, 0

	for i := len(garbage) - 1; i > 0; i-- {
		volume += len(garbage[i])

		if mIndex == 0 || pIndex == 0 || gIndex == 0 {
			mIndex, pIndex, gIndex = detector(garbage[i], i, mIndex, pIndex, gIndex)
		}
	}

	for i := 1; i < len(travel); i++ {
		travel[i] += travel[i-1]
	}

	return volume + distance(travel, mIndex) + distance(travel, pIndex) + distance(travel, gIndex)
}

func distance(travel []int, house int) int {
	if house == 0 {
		return 0
	}

	return travel[house-1]
}

func detector(garbage string, house, mIndex, pIndex, gIndex int) (int, int, int) {
	for i := 0; i < len(garbage) && (mIndex == 0 || pIndex == 0 || gIndex == 0); i++ {
		switch {
		case garbage[i] == 'M' && mIndex == 0:
			mIndex = house
		case garbage[i] == 'P' && pIndex == 0:
			pIndex = house
		case garbage[i] == 'G' && gIndex == 0:
			gIndex = house
		}
	}

	return mIndex, pIndex, gIndex
}

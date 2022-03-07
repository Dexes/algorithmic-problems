package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	result := 0
	for i := 0; i < len(boxTypes); i++ {
		result += boxTypes[i][0] * boxTypes[i][1]
		truckSize -= boxTypes[i][0]
		if truckSize <= 0 {
			result += truckSize * boxTypes[i][1]
			break
		}
	}

	return result
}

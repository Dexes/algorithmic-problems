package main

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	wIndex := 0
	for i1, i2, i3 := 0, 0, 0; i1 < len(arr1) && i2 < len(arr2) && i3 < len(arr3); {
		switch {
		case arr1[i1] == arr2[i2] && arr1[i1] == arr3[i3]:
			arr1[wIndex], wIndex = arr1[i1], wIndex+1
			i1, i2, i3 = i1+1, i2+1, i3+1
		case arr1[i1] < arr2[i2]:
			i1++
		case arr2[i2] < arr3[i3]:
			i2++
		default:
			i3++
		}
	}

	return arr1[:wIndex]
}

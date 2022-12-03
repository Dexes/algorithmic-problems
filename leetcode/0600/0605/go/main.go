package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	n -= processBorders(flowerbed)
	if n <= 0 {
		return true
	}

	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i+1] == 1 {
			i += 2
			continue
		}

		if flowerbed[i] == 1 {
			i++
			continue
		}

		if flowerbed[i-1] == 1 {
			continue
		}

		i, n = i+1, n-1
		if n == 0 {
			return true
		}
	}

	return false
}

func processBorders(flowerbed []int) int {
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return 1
	}

	result := 0
	if len(flowerbed) > 1 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		result++
	}

	if len(flowerbed) > 2 && flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		flowerbed[len(flowerbed)-1] = 1
		result++
	}

	return result
}

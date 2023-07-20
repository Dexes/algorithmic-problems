package main

func asteroidCollision(asteroids []int) []int {
	stackIndex := -1
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stackIndex++
			asteroids[stackIndex] = asteroids[i]
			continue
		}

		size := -asteroids[i]
		for ; stackIndex >= 0 && asteroids[stackIndex] > 0 && asteroids[stackIndex] < size; stackIndex-- {
		}

		if stackIndex < 0 || asteroids[stackIndex] < 0 {
			stackIndex++
			asteroids[stackIndex] = asteroids[i]
			continue
		}

		if asteroids[stackIndex] == size {
			stackIndex--
		}
	}

	return asteroids[:stackIndex+1]
}

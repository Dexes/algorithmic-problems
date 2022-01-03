package main

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > mass {
			return false
		}

		mass += asteroids[i]
	}

	return true
}

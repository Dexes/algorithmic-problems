package main

import (
	"sort"
)

func closestPrimes(left, right int) []int {
	left, right = toIndices(left, right)
	if left >= right {
		return emptyResult
	}

	first, second := primes[left], primes[left+1]
	diff := second - first
	for i := left + 2; i <= right; i++ {
		if x := primes[i] - primes[i-1]; x < diff {
			first, second, diff = primes[i-1], primes[i], x
		}
	}

	return []int{first, second}
}

func toIndices(left, right int) (int, int) {
	leftIndex := sort.Search(len(primes), func(i int) bool { return primes[i] >= left })
	rightIndex := sort.Search(len(primes)-leftIndex, func(i int) bool { return primes[i+leftIndex] >= right }) + leftIndex

	if primes[leftIndex] < left {
		leftIndex++
	}

	if rightIndex == len(primes) || primes[rightIndex] > right {
		rightIndex--
	}

	return leftIndex, rightIndex
}

var (
	primes      []int
	emptyResult = []int{-1, -1}
)

func init() {
	primes = make([]int, 0, 78498)

	set := make([]bool, 1e6+1)
	for num := 2; num < len(set); num++ {
		if set[num] {
			continue
		}

		for x := num * num; x < len(set); x += num {
			set[x] = true
		}

		primes = append(primes, num)
	}
}

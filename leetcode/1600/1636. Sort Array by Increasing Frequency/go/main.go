package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
}

type Element struct {
	Number    int
	Frequency int
}

func frequencySort(nums []int) []int {
	elements := count(nums)
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].Frequency == elements[j].Frequency {
			return elements[i].Number > elements[j].Number
		}

		return elements[i].Frequency < elements[j].Frequency
	})

	for i, j := 0, 0; i < len(elements); i++ {
		for ; elements[i].Frequency > 0; elements[i].Frequency-- {
			nums[j], j = elements[i].Number, j+1
		}
	}

	return nums
}

func count(nums []int) []Element {
	result := make([]Element, 201)
	for i := 0; i < len(nums); i++ {
		result[nums[i]+100].Frequency++
		result[nums[i]+100].Number = nums[i]
	}

	return result
}

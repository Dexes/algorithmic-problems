package main

func twoSum(numbers []int, target int) []int {
	for i := 0; ; i++ {
		number := target - numbers[i]
		left, right := i+1, len(numbers)
		for left < right {
			middle := left + (right-left)/2
			if number < numbers[middle] {
				right = middle
			} else if number > numbers[middle] {
				left = middle + 1
			} else {
				return []int{i + 1, middle + 1}
			}
		}
	}
}

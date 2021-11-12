package main

func rotate(nums []int, k int) {
	length := len(nums)
	processed, index := 0, -1

	for processed < length {
		index++
		destination := (index + k) % length
		buffer := nums[index]

		for {
			nums[destination], buffer = buffer, nums[destination]
			processed++
			if destination == index {
				break
			}

			destination = (destination + k) % length
		}
	}
}

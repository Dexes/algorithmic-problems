package main

func sortColors(nums []int) {
	counter := []int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	for i, j := 0, 0; i < len(counter); i++ {
		for ; counter[i] > 0; j, counter[i] = j+1, counter[i]-1 {
			nums[j] = i
		}
	}
}

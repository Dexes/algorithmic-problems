package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	data := make([][]byte, 101)
	fillNumbers(data, nums1, 0)
	fillNumbers(data, nums2, 1)
	fillNumbers(data, nums3, 2)

	result := make([]int, 0, 101)
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			continue
		}

		if data[i][0]+data[i][1]+data[i][2] > 1 {
			result = append(result, i)
		}
	}

	return result
}

func fillNumbers(data [][]byte, nums []int, index int) {
	for i := 0; i < len(nums); i++ {
		if data[nums[i]] == nil {
			data[nums[i]] = make([]byte, 3)
		}

		data[nums[i]][index] = 1
	}
}

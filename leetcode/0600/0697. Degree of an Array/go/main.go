package main

type Frequency struct {
	Count       int
	Left, Right int
}

func findShortestSubArray(nums []int) int {
	result, count := len(nums), 0
	for _, freq := range frequencies(nums) {
		if freq.Count > count {
			result, count = freq.Right-freq.Left, freq.Count
			continue
		}

		if freq.Count == count && freq.Right-freq.Left < result {
			result = freq.Right - freq.Left
		}
	}

	return result + 1
}

func frequencies(nums []int) map[int]*Frequency {
	result := make(map[int]*Frequency)
	for i := 0; i < len(nums); i++ {
		if x, ok := result[nums[i]]; ok {
			x.Count++
			x.Right = i
		} else {
			result[nums[i]] = &Frequency{Count: 1, Left: i, Right: i}
		}
	}

	return result
}

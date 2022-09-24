package main

func findSubarrays(nums []int) bool {
	hash := make(map[int]struct{})
	for i := 1; i < len(nums); i++ {
		x := nums[i] + nums[i-1]
		if _, ok := hash[x]; ok {
			return true
		}

		hash[x] = struct{}{}
	}

	return false
}

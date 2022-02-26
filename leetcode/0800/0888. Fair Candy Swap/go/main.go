package main

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	avg, bob := (sum(bobSizes)-sum(aliceSizes))/2, toSet(bobSizes)
	for i := 0; ; i++ {
		// avg + aliceSizes[i] == ((bobSum + aliceSizes[i]) - (aliceSum - aliceSizes[i])) / 2
		if _, ok := bob[avg+aliceSizes[i]]; ok {
			return []int{aliceSizes[i], avg + aliceSizes[i]}
		}
	}
}

func sum(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}

	return result
}

func toSet(nums []int) map[int]struct{} {
	result := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		result[nums[i]] = struct{}{}
	}

	return result
}

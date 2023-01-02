package main

func longestConsecutive(nums []int) (result int) {
	if len(nums) < 2 {
		return len(nums)
	}

	set := toSet(nums)
	for num := range set {
		if l := length(num, set); l > result {
			result = l
		}
	}

	return result
}

func length(num int, set map[int]int) int {
	result := 1

	for x := num + 1; ; x++ {
		switch set[x] {
		case -1:
			return -1
		case 0:
			set[num] = result
			return result
		case 1:
			set[x], result = -1, result+1
		default:
			result += set[x]
			set[num] = result
			return result
		}
	}
}

func toSet(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num] = 1
	}

	return result
}

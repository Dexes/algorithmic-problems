package main

func separateDigits(nums []int) []int {
	result := make([]int, 0, len(nums)<<1)
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] < 10:
			result = append(result, nums[i])

		case nums[i] < 100:
			result = append(result, nums[i]/10, nums[i]%10)

		default:
			j := len(result)
			for ; nums[i] > 0; nums[i] /= 10 {
				result = append(result, nums[i]%10)
			}

			reverse(result[j:])
		}
	}

	return result
}

func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

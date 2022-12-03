package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var ls, rs int
	result := make([][]int, 0, 100)

	sort.Ints(nums)

	for i := len(nums) - 1; i >= 3; i-- {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			continue
		}

		for j := i - 1; j >= 2; j-- {
			if j+1 < i && nums[j] == nums[j+1] {
				continue
			}

			sum := nums[i] + nums[j]

			for left, right := 0, j-1; left < right; {
				switch x := sum + nums[left] + nums[right]; {
				case x < target:
					left++
				case x > target:
					right--
				default:
					result = append(result, []int{nums[left], nums[right], nums[j], nums[i]})

					for ls, left = left, left+1; left < right && nums[ls] == nums[left]; left++ {
					}

					for rs, right = right, right-1; left < right && nums[rs] == nums[right]; right-- {
					}
				}
			}
		}
	}

	return result
}

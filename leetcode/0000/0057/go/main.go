package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if result, success := edgeInsert(intervals, newInterval); success {
		return result
	}

	left, right := 0, 0
	for ; intervals[left][1] < newInterval[0]; left++ {
	}

	if newInterval[1] < intervals[left][0] {
		result := append(intervals[:left+1], intervals[left:]...)
		result[left] = newInterval
		return result
	}

	for right = left; right < len(intervals) && intervals[right][0] <= newInterval[1]; right++ {
	}

	intervals[left][0] = min(intervals[left][0], newInterval[0])
	intervals[left][1] = max(intervals[right-1][1], newInterval[1])

	return append(intervals[:left+1], intervals[right:]...)
}

func edgeInsert(intervals [][]int, newInterval []int) ([][]int, bool) {
	if len(intervals) == 0 {
		return [][]int{newInterval}, true
	}

	if newInterval[1] < intervals[0][0] {
		result := make([][]int, 0, len(intervals)+1)
		result = append(result, newInterval)
		return append(result, intervals...), true
	}

	lastIndex := len(intervals) - 1
	if intervals[lastIndex][1] < newInterval[0] {
		result := make([][]int, 0, len(intervals)+1)
		result = append(result, intervals...)
		return append(result, newInterval), true
	}

	return nil, false
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

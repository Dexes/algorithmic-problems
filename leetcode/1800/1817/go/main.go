package main

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	var (
		times  map[int]struct{}
		users  = make(map[int]map[int]struct{})
		result = make([]int, k)
	)

	for _, log := range logs {
		if times = users[log[0]]; times == nil {
			times = make(map[int]struct{})
			users[log[0]] = times
		}

		times[log[1]] = struct{}{}
	}

	for _, times = range users {
		result[len(times)-1]++
	}

	return result
}

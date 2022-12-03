package main

type Point struct {
	N    int
	Time int
}

func networkDelayTime(routes [][]int, n int, k int) int {
	graph, queue, times := toGraph(routes, n), make(map[int]struct{}), makeTimes(n)
	queue[k-1], times[k-1] = struct{}{}, 0

	for len(queue) > 0 {
		n = pop(queue, times)
		for _, t := range graph[n] {
			if times[t.N] == -1 {
				times[t.N] = times[n] + t.Time
				queue[t.N] = struct{}{}
				continue
			}

			if times[t.N] > times[n]+t.Time {
				times[t.N] = times[n] + t.Time
			}
		}
	}

	return max(times)
}

func pop(queue map[int]struct{}, times []int) int {
	result, time := 0, -1
	for x := range queue {
		if time == -1 || times[x] < time {
			result, time = x, times[x]
		}
	}

	delete(queue, result)
	return result
}

func max(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			return -1
		}

		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}

func toGraph(routes [][]int, n int) [][]Point {
	result := make([][]Point, n)
	for i := 0; i < len(routes); i++ {
		from, to := routes[i][0]-1, routes[i][1]-1
		if cap(result[from]) == 0 {
			result[from] = make([]Point, 0, 10)
		}

		result[from] = append(result[from], Point{N: to, Time: routes[i][2]})
	}

	return result
}

func makeTimes(n int) []int {
	result := make([]int, n)
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}

	return result
}

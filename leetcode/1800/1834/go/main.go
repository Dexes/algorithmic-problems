package main

import "sort"

func getOrder(input [][]int) []int {
	tasks, result := toTasks(input), make([]int, len(input))
	heap, time := tasks[:0], 0

	for i := 0; i < len(result); i++ {
		if len(heap) == 0 && time < tasks[0].EnqueueTime {
			time = tasks[0].EnqueueTime
		}

		if len(tasks) > 0 {
			heap, tasks = shiftTasks(time, heap, tasks)
		}

		result[i] = heap[0].Index
		time += heap[0].ProcessingTime
		heap = removeTask(heap)
	}

	return result
}

func toTasks(input [][]int) []Task {
	result := make([]Task, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = Task{Index: i, EnqueueTime: input[i][0], ProcessingTime: input[i][1]}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].EnqueueTime < result[j].EnqueueTime
	})

	return result
}

func shiftTasks(time int, heap, tasks []Task) ([]Task, []Task) {
	firstUnavailableTask := sort.Search(len(tasks), func(i int) bool {
		return tasks[i].EnqueueTime > time
	})

	for i := 0; i < firstUnavailableTask; i++ {
		heap = append(heap, tasks[i])
		siftUp(heap, len(heap)-1)
	}

	return heap, tasks[firstUnavailableTask:]
}

func removeTask(heap []Task) []Task {
	lastIndex := len(heap) - 1
	heap[0], heap = heap[lastIndex], heap[:lastIndex]
	siftDown(heap, 0)

	return heap
}

func siftDown(heap []Task, i int) {
	for cIndex := getChildIndex(heap, i); cIndex > 0 && less(heap[i], heap[cIndex]); cIndex = getChildIndex(heap, i) {
		heap[i], heap[cIndex] = heap[cIndex], heap[i]
		i = cIndex
	}
}

func siftUp(heap []Task, i int) {
	for pIndex := (i - 1) / 2; less(heap[pIndex], heap[i]); pIndex = (i - 1) / 2 {
		heap[i], heap[pIndex] = heap[pIndex], heap[i]
		i = pIndex
	}
}

func getChildIndex(heap []Task, i int) int {
	switch lIndex, rIndex := i*2+1, i*2+2; {
	case lIndex >= len(heap):
		return -1
	case rIndex >= len(heap):
		return lIndex
	case less(heap[rIndex], heap[lIndex]):
		return lIndex
	default:
		return rIndex
	}
}

func less(x, y Task) bool {
	if x.ProcessingTime == y.ProcessingTime {
		return x.Index > y.Index
	}

	return x.ProcessingTime > y.ProcessingTime
}

type Task struct {
	Index          int
	EnqueueTime    int
	ProcessingTime int
}

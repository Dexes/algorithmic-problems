package main

type MinStack struct {
	data     []int
	minimums []int
	index    int
}

func Constructor() MinStack {
	return MinStack{data: make([]int, 10000), minimums: make([]int, 10000), index: -1}
}

func (s *MinStack) Push(val int) {
	s.index++

	s.data[s.index] = val
	if s.index == 0 || s.minimums[s.index-1] > val {
		s.minimums[s.index] = val
	} else {
		s.minimums[s.index] = s.minimums[s.index-1]
	}
}

func (s *MinStack) Pop() {
	s.index--
}

func (s *MinStack) Top() int {
	return s.data[s.index]
}

func (s *MinStack) GetMin() int {
	return s.minimums[s.index]
}

package main

type TwoSum struct {
	counter map[int]int
}

func Constructor() TwoSum {
	return TwoSum{counter: make(map[int]int)}
}

func (x TwoSum) Add(number int) {
	x.counter[number]++
}

func (x TwoSum) Find(value int) bool {
	for first, count := range x.counter {
		second := value - first
		if _, ok := x.counter[second]; ok {
			if first != second || count > 1 {
				return true
			}
		}
	}

	return false
}

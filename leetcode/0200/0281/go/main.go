package main

type ZigzagIterator struct {
	first, second []int
}

func Constructor(first, second []int) *ZigzagIterator {
	if len(first) == 0 {
		first, second = second, first
	}

	return &ZigzagIterator{first: first, second: second}
}

func (x *ZigzagIterator) next() (result int) {
	result, x.first = x.first[0], x.first[1:]
	if len(x.second) > 0 {
		x.first, x.second = x.second, x.first
	}

	return result
}

func (x *ZigzagIterator) hasNext() bool {
	return len(x.first) > 0
}

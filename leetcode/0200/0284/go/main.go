package main

type Iterator struct{}

func (x *Iterator) hasNext() bool { return true }
func (x *Iterator) next() int     { return 0 }

type PeekingIterator struct {
	iterator     *Iterator
	hasNextValue bool
	nextValue    int
}

func Constructor(iterator *Iterator) *PeekingIterator {
	result := &PeekingIterator{iterator: iterator, hasNextValue: iterator.hasNext(), nextValue: 0}
	if result.hasNextValue {
		result.nextValue = iterator.next()
	}

	return result
}

func (x *PeekingIterator) hasNext() bool {
	return x.hasNextValue
}

func (x *PeekingIterator) next() (result int) {
	result = x.nextValue

	if x.iterator.hasNext() {
		x.nextValue = x.iterator.next()
	} else {
		x.hasNextValue = false
	}

	return result
}

func (x *PeekingIterator) peek() int {
	return x.nextValue
}

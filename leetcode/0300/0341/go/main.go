package main

type NestedIterator struct {
	stack      []*NestedInteger
	stackIndex int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	result := &NestedIterator{stack: make([]*NestedInteger, 1000), stackIndex: -1}
	result.push(nestedList)
	result.flatten()

	return result
}

func (x *NestedIterator) Next() int {
	return x.pop().GetInteger()
}

func (x *NestedIterator) HasNext() bool {
	x.flatten()
	return x.stackIndex >= 0 && x.stack[x.stackIndex].IsInteger()
}

func (x *NestedIterator) pop() *NestedInteger {
	result := x.stack[x.stackIndex]
	x.stackIndex--

	return result
}

func (x *NestedIterator) push(nestedList []*NestedInteger) {
	for i := len(nestedList) - 1; i >= 0; i-- {
		x.stackIndex++
		x.stack[x.stackIndex] = nestedList[i]
	}
}

func (x *NestedIterator) flatten() {
	for x.stackIndex >= 0 && !x.stack[x.stackIndex].IsInteger() {
		x.push(x.pop().GetList())
	}
}

type NestedInteger struct {
}

func (x NestedInteger) IsInteger() bool {
	return true
}

func (x NestedInteger) GetInteger() int {
	return 0
}

func (x *NestedInteger) SetInteger(value int) {
}

func (x *NestedInteger) Add(elem NestedInteger) {
}

func (x NestedInteger) GetList() []*NestedInteger {
	return nil
}

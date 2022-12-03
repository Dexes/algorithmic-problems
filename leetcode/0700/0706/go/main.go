package main

type MyHashMap struct {
	data []*BucketNode // replace with RB tree
}

func Constructor() MyHashMap {
	return MyHashMap{data: make([]*BucketNode, 1000)}
}

func (x *MyHashMap) Put(key int, value int) {
	bucket, _, current := x.findPrevAndCurrent(key)

	if current != nil {
		current.Value = value
	} else {
		x.data[bucket] = &BucketNode{Key: key, Value: value, Next: x.data[bucket]}
	}
}

func (x *MyHashMap) Get(key int) int {
	_, _, current := x.findPrevAndCurrent(key)
	if current != nil {
		return current.Value
	}

	return -1
}

func (x *MyHashMap) Remove(key int) {
	bucket, prev, current := x.findPrevAndCurrent(key)
	if current == nil {
		return
	}

	if prev == nil {
		x.data[bucket] = current.Next
		return
	}

	prev.Next = current.Next
}

func (x *MyHashMap) findPrevAndCurrent(key int) (bucket int, prev, current *BucketNode) {
	bucket = x.getBucket(key)
	prev, current = nil, x.data[bucket]
	for current != nil {
		if current.Key == key {
			return bucket, prev, current
		}

		prev, current = current, current.Next
	}

	return bucket, nil, nil
}

func (x *MyHashMap) getBucket(key int) int {
	if key == 0 {
		return 0
	}

	return key % len(x.data)
}

type BucketNode struct {
	Key   int
	Value int
	Next  *BucketNode
}

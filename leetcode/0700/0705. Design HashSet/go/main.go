package main

type MyHashSet struct {
	data []*BucketNode // replace with RB tree
}

func Constructor() MyHashSet {
	return MyHashSet{data: make([]*BucketNode, 1000)}
}

func (x *MyHashSet) Add(key int) {
	bucket, _, current := x.findPrevAndCurrent(key)
	if current == nil {
		x.data[bucket] = &BucketNode{Key: key, Next: x.data[bucket]}
	}
}

func (x *MyHashSet) Remove(key int) {
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

func (x *MyHashSet) Contains(key int) bool {
	_, _, current := x.findPrevAndCurrent(key)
	return current != nil
}

func (x *MyHashSet) findPrevAndCurrent(key int) (bucket int, prev, current *BucketNode) {
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

func (x *MyHashSet) getBucket(key int) int {
	if key == 0 {
		return 0
	}

	return key % len(x.data)
}

type BucketNode struct {
	Key  int
	Next *BucketNode
}

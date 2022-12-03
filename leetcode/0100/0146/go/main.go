package main

type LRUCache struct {
	capacity int
	data     map[int]*CacheNode
	head     *CacheNode
	tail     *CacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*CacheNode, capacity),
	}
}

func (x *LRUCache) Get(key int) int {
	var (
		node   *CacheNode
		exists bool
	)

	if node, exists = x.data[key]; !exists {
		return -1
	}

	x.toTail(node)

	return x.tail.Val
}

func (x *LRUCache) Put(key int, value int) {
	if len(x.data) == 0 {
		node := &CacheNode{Key: key, Val: value}
		x.head, x.tail = node, node
		x.data[key] = node
		return
	}

	if node, exists := x.data[key]; exists {
		node.Val = value
		x.toTail(node)
		return
	}

	node := &CacheNode{Key: key, Val: value}
	x.tail.Next, x.tail, x.data[key] = node, node, node

	if len(x.data) > x.capacity {
		delete(x.data, x.head.Key)
		x.head = x.head.Next
	}
}

func (x *LRUCache) toTail(node *CacheNode) {
	if node == x.tail {
		return
	}

	tail := &CacheNode{Key: node.Key, Val: node.Val}
	x.tail.Next, x.tail, x.data[tail.Key] = tail, tail, tail

	node.Key, node.Val, node.Next = node.Next.Key, node.Next.Val, node.Next.Next
	x.data[node.Key] = node
}

type CacheNode struct {
	Key, Val int
	Next     *CacheNode
}

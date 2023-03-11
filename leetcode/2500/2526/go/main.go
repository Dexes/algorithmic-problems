package main

type DataStream struct {
	value, k int
	size     int
}

func Constructor(value int, k int) DataStream {
	return DataStream{value: value, k: k, size: 0}
}

func (x *DataStream) Consec(num int) bool {
	if num != x.value {
		x.size = 0
		return false
	}

	x.size++
	return x.size >= x.k
}

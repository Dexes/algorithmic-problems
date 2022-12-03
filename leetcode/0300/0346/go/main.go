package main

type MovingAverage struct {
	data  []int
	sum   int
	index int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{data: make([]int, size), sum: 0, index: 0}
}

func (x *MovingAverage) Next(val int) float64 {
	index := x.index % len(x.data)
	x.sum += val - x.data[index]
	x.data[index] = val
	x.index++

	return float64(x.sum) / float64(min(x.index, len(x.data)))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

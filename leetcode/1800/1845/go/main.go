package main

type SeatManager struct {
	seats []int
}

func Constructor(n int) SeatManager {
	seats := make([]int, n)
	for i := 0; i < n; i++ {
		seats[i] = i + 1
	}

	return SeatManager{seats: seats}
}

func (x *SeatManager) Reserve() int {
	result := x.seats[0]

	lastIndex := len(x.seats) - 1
	x.seats[0], x.seats = x.seats[lastIndex], x.seats[:lastIndex]
	x.siftDown(0)

	return result
}

func (x *SeatManager) Unreserve(seatNumber int) {
	x.seats = append(x.seats, seatNumber)
	x.siftUp(len(x.seats) - 1)
}

func (x *SeatManager) siftUp(i int) {
	for pIndex := (i - 1) / 2; x.less(pIndex, i); pIndex = (i - 1) / 2 {
		x.swap(i, pIndex)
		i = pIndex
	}
}

func (x *SeatManager) siftDown(i int) {
	for cIndex := x.getChildIndex(i); cIndex > 0 && x.less(i, cIndex); cIndex = x.getChildIndex(i) {
		x.swap(i, cIndex)
		i = cIndex
	}
}

func (x *SeatManager) getChildIndex(i int) int {
	lIndex, rIndex := i*2+1, i*2+2
	if lIndex >= len(x.seats) {
		return -1
	}

	if rIndex >= len(x.seats) {
		return lIndex
	}

	if x.less(rIndex, lIndex) {
		return lIndex
	}

	return rIndex
}

func (x *SeatManager) swap(i, j int) {
	x.seats[i], x.seats[j] = x.seats[j], x.seats[i]
}

func (x *SeatManager) less(a, b int) bool {
	return x.seats[a] > x.seats[b]
}

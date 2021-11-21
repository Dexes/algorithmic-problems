package main

type ParkingSystem struct {
	lots []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{lots: []int{big, medium, small}}
}

func (x *ParkingSystem) AddCar(carType int) bool {
	if x.lots[carType-1] == 0 {
		return false
	}

	x.lots[carType-1]--
	return true
}

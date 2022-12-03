package main

type ParkingSystem struct {
	lots [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{lots: [3]int{big, medium, small}}
}

func (x *ParkingSystem) AddCar(carType int) bool {
	if carType--; x.lots[carType] == 0 {
		return false
	}

	x.lots[carType]--
	return true
}

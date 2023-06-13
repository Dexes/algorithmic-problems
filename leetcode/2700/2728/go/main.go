package main

type Street interface {
	OpenDoor()
	CloseDoor()
	IsDoorOpen() bool
	MoveRight()
	MoveLeft()
}

func houseCount(street Street, k int) (result int) {
	for i := 0; i < k; i++ {
		street.OpenDoor()
		street.MoveRight()
	}

	street.CloseDoor()
	street.MoveRight()

	for result = 1; street.IsDoorOpen(); result++ {
		street.MoveRight()
	}

	return result
}

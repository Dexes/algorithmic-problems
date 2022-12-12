package main

type PhoneDirectory struct {
	index   int
	numbers []bool
}

func Constructor(capacity int) PhoneDirectory {
	return PhoneDirectory{index: 0, numbers: make([]bool, capacity)}
}

func (d *PhoneDirectory) Get() int {
	for ; d.index < len(d.numbers) && d.numbers[d.index]; d.index++ {
	}

	if d.index == len(d.numbers) {
		return -1
	}

	d.numbers[d.index] = true
	return d.index
}

func (d *PhoneDirectory) Check(number int) bool {
	return !d.numbers[number]
}

func (d *PhoneDirectory) Release(number int) {
	if number < d.index {
		d.index = number
	}

	d.numbers[number] = false
}

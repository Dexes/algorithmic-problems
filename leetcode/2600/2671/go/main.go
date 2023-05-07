package main

const max = 1e5 + 1

type FrequencyTracker struct {
	numbers     []int
	frequencies []int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{numbers: make([]int, max), frequencies: make([]int, max)}
}

func (t FrequencyTracker) Add(number int) {
	t.frequencies[t.numbers[number]]--
	t.numbers[number]++
	t.frequencies[t.numbers[number]]++
}

func (t FrequencyTracker) DeleteOne(number int) {
	if t.numbers[number] == 0 {
		return
	}

	t.frequencies[t.numbers[number]]--
	t.numbers[number]--
	t.frequencies[t.numbers[number]]++
}

func (t FrequencyTracker) HasFrequency(frequency int) bool {
	return t.frequencies[frequency] > 0
}

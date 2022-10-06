package main

type TimeMap map[string][]Item

func Constructor() TimeMap {
	return make(TimeMap)
}

func (m TimeMap) Set(key string, value string, timestamp int) {
	x := Item{timestamp: timestamp, value: value}

	if items, ok := m[key]; ok {
		m[key] = append(items, x)
		return
	}

	m[key] = append(make([]Item, 0, 8), x)
}

func (m TimeMap) Get(key string, timestamp int) string {
	if items, ok := m[key]; ok {
		return search(items, timestamp)
	}

	return ""
}

func search(x []Item, t int) string {
	left, right := 0, len(x)-1
	for left < right {
		if middle := left + (right-left)/2 + 1; x[middle].timestamp > t {
			right = middle - 1
		} else {
			left = middle
		}
	}

	if result := x[left]; result.timestamp <= t {
		return result.value
	}

	return ""
}

type Item struct {
	timestamp int
	value     string
}

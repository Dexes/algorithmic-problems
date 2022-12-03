package main

type OrderedStream struct {
	current int
	parts   []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		current: 0,
		parts:   make([]string, n),
	}
}

func (x *OrderedStream) Insert(i int, value string) []string {
	i--
	x.parts[i] = value

	if i != x.current {
		return []string{}
	}

	result := make([]string, 0, len(x.parts)-i)
	for ; i < len(x.parts) && x.parts[i] != ""; i++ {
		result = append(result, x.parts[i])
	}

	x.current = i
	return result
}

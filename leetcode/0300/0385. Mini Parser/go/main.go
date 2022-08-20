package main

func deserialize(s string) *NestedInteger {
	result, _ := readNested(s, 0)
	return &result
}

func readNested(s string, i int) (NestedInteger, int) {
	if s[i] == '-' || '0' <= s[i] && s[i] <= '9' {
		return readInt(s, i)
	}

	var (
		result = NestedInteger{}
		x      NestedInteger
	)

	for i++; s[i] != ']'; i++ {
		if s[i] == ',' {
			continue
		}

		x, i = readNested(s, i)
		result.Add(x)
	}

	return result, i
}

func readInt(s string, i int) (NestedInteger, int) {
	num, sign := 0, 1
	if s[i] == '-' {
		sign = -1
		i++
	}

	for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		num = num*10 + int(s[i]-'0')
	}

	result := NestedInteger{}
	result.SetInteger(sign * num)

	return result, i - 1
}

type NestedInteger struct {
	NestedInterface
}

type NestedInterface interface {
	IsInteger() bool
	GetInteger() int
	SetInteger(value int)
	Add(elem NestedInteger)
	GetList() []*NestedInteger
}

package main

type key struct {
	data             []byte
	index            int
	currentGroupSize int
	groupSize        int
}

func (x *key) getValidLetter(letter byte) byte {
	if letter >= 'a' && letter <= 'z' {
		return letter - 32
	}

	return letter
}

func (x *key) append(letter byte) {
	x.data[x.index] = x.getValidLetter(letter)
	x.index--
	x.currentGroupSize++

	if x.currentGroupSize == x.groupSize && x.index >= 0 {
		x.data[x.index] = '-'
		x.index--
		x.currentGroupSize = 0
	}
}

func (x *key) toString() string {
	if x.index == len(x.data)-1 {
		return ""
	}

	if x.data[x.index+1] == '-' {
		return string(x.data[x.index+2:])
	}

	return string(x.data[x.index+1:])
}

func makeKey(groupSize, cap int) *key {
	return &key{data: make([]byte, cap), index: cap - 1, currentGroupSize: 0, groupSize: groupSize}
}

func licenseKeyFormatting(s string, k int) string {
	result := makeKey(k, len(s)*2)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}

		result.append(s[i])
	}

	return result.toString()
}

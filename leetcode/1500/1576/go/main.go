package main

func modifyString(s string) string {
	b, lastIndex := []byte(s), len(s)-1
	for i := 1; i < lastIndex; i++ {
		if b[i] == '?' {
			b[i] = makeChar(b, i)
		}
	}

	b[0] = makeFirstChar(b)
	b[lastIndex] = makeLastChar(b)

	return string(b)
}

func makeChar(b []byte, index int) byte {
	aExists := b[index-1] == 'a' || b[index+1] == 'a'
	bExists := b[index-1] == 'b' || b[index+1] == 'b'

	switch {
	case aExists && bExists:
		return 'c'
	case aExists:
		return 'b'
	default:
		return 'a'
	}
}

func makeFirstChar(b []byte) byte {
	switch {
	case b[0] != '?':
		return b[0]
	case len(b) == 1:
		return 'a'
	case b[1] == 'a':
		return 'b'
	default:
		return 'a'
	}
}

func makeLastChar(b []byte) byte {
	switch lastIndex := len(b) - 1; {
	case b[lastIndex] != '?':
		return b[lastIndex]
	case len(b) == 1:
		return 'a'
	case b[lastIndex-1] == 'a':
		return 'b'
	default:
		return 'a'
	}
}

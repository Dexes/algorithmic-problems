package main

func lengthLongestPath(input string) (result int) {
	stack, stackIndex, pathLength := make([]FSObject, 100), -1, 0
	for i := 0; i < len(input); {
		object := readObject(input, i)
		i += object.Depth + object.NameLength + 1

		for ; stackIndex >= 0 && stack[stackIndex].Depth >= object.Depth; stackIndex-- {
			pathLength -= stack[stackIndex].NameLength
		}

		stackIndex++
		stack[stackIndex] = object
		pathLength += object.NameLength

		if object.IsFile {
			result = max(result, pathLength+object.Depth)
		}
	}

	return result
}

func readObject(s string, startIndex int) (result FSObject) {
	i := startIndex
	for ; s[i] == '\t'; i++ {
	}
	result.Depth = i - startIndex

	startIndex = i
	for ; i < len(s) && s[i] != '\n'; i++ {
		if s[i] == '.' {
			result.IsFile = true
		}
	}
	result.NameLength = i - startIndex

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type FSObject struct {
	Depth      int
	NameLength int
	IsFile     bool
}

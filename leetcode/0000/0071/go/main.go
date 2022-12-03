package main

func simplifyPath(path string) string {
	result, resultIndex := make([]byte, len(path)), 0
	foldersIndices, foldersIndex := make([]int, 500), 0

	for i, j := 0, 0; i < len(path); i++ {
		for j = i; j < len(path) && path[j] != '/'; j++ {
		}

		switch length := j - i; {
		case length == 0 || length == 1 && path[i] == '.':
			continue

		case length == 2 && path[i] == '.' && path[i+1] == '.':
			if foldersIndex > 1 {
				resultIndex, foldersIndex = foldersIndices[foldersIndex-1], foldersIndex-2
			} else {
				resultIndex, foldersIndex = 0, 0
			}

		default:
			result[resultIndex] = '/'
			for resultIndex++; i < j; i++ {
				result[resultIndex] = path[i]
				resultIndex++
			}

			foldersIndices[foldersIndex] = resultIndex
			foldersIndex++
		}
	}

	if resultIndex == 0 {
		return "/"
	}

	return string(result[:resultIndex])
}

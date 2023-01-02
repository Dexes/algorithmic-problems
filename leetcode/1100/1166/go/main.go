package main

type FileSystem map[string]int

func Constructor() FileSystem {
	result := make(FileSystem)
	result[""] = 0

	return result
}

func (x FileSystem) CreatePath(path string, value int) bool {
	if _, ok := x[path]; ok {
		return false
	}

	if _, ok := x[parent(path)]; !ok {
		return false
	}

	x[path] = value
	return true
}

func (x FileSystem) Get(path string) int {
	if val, ok := x[path]; ok {
		return val
	}

	return -1
}

func parent(path string) string {
	index := len(path) - 1
	for ; path[index] != '/'; index-- {
	}

	return path[:index]
}

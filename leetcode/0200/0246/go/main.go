package main

var reversed = []byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}

func isStrobogrammatic(num string) bool {
	for left, right := 0, len(num)-1; left <= right; left, right = left+1, right-1 {
		if reversed[num[left]] != num[right] {
			return false
		}
	}

	return true
}

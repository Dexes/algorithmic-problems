package main

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i]&1 == 1 {
			return num[:i+1]
		}
	}

	return ""
}

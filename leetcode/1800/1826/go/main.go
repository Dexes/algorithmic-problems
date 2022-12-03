package main

func badSensor(s1 []int, s2 []int) int {
	i := 0
	for ; i < len(s1) && s1[i] == s2[i]; i++ {
	}

	for i++; i < len(s1) && s1[i] == s2[i-1] && s1[i-1] == s2[i]; i++ {
	}

	if i >= len(s1) {
		return -1
	}

	if s1[i-1] == s2[i] {
		return 1
	}

	return 2
}

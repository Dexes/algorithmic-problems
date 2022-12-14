package main

func dividePlayers(skill []int) int64 {
	counter, sum := make(map[int]int), 0
	for i := 0; i < len(skill); i++ {
		counter[skill[i]]++
		sum += skill[i]
	}

	half := len(skill) / 2
	chemistry, mod := sum/half, sum%half

	if mod > 0 {
		return -1
	}

	result := int64(0)
	for a, num := range counter {
		switch b := chemistry - a; {
		case num == 0:
			continue
		case a > chemistry || counter[b] != num || (a == b && num%2 == 1):
			return -1
		case a == b:
			result += int64(a * b * num / 2)
		default:
			result += int64(a * b * num)
			counter[b] = 0
		}
	}

	return result
}

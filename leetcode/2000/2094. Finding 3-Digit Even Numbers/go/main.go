package main

func findEvenNumbers(numbers []int) []int {
	digits := make([]byte, 10)
	for i := 0; i < len(numbers); i++ {
		digits[numbers[i]]++
	}

	result := make([]int, 0, 900)
	for first := 1; first <= 9; first += 1 {
		if digits[first] == 0 {
			continue
		}

		digits[first]--
		for second := 0; second <= 9; second += 1 {
			if digits[second] == 0 {
				continue
			}

			digits[second]--
			for third := 0; third <= 8; third += 2 {
				if digits[third] > 0 {
					result = append(result, first*100+second*10+third)
				}
			}
			digits[second]++
		}
		digits[first]++
	}

	return result
}

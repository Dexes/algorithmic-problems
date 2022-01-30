package main

func findRestaurant(list1 []string, list2 []string) []string {
	restaurants, index, min := toMap(list2), 0, len(list1)+len(list2)
	for i := 0; i < len(list1) && i <= min; i++ {
		j, exists := restaurants[list1[i]]
		if !exists {
			continue
		}

		if i+j < min {
			list1[0], index, min = list1[i], 1, i+j
			continue
		}

		if i+j == min {
			list1[index], index = list1[i], index+1
		}
	}

	return list1[:index]
}

func toMap(list []string) map[string]int {
	result := make(map[string]int)
	for i := 0; i < len(list); i++ {
		result[list[i]] = i
	}

	return result
}

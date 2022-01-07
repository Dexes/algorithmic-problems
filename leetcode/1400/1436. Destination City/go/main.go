package main

func destCity(paths [][]string) string {
	from, to := toMap(paths, 0), toMap(paths, 1)
	for city, _ := range to {
		if !from[city] {
			return city
		}
	}

	return ""
}

func toMap(paths [][]string, index int) map[string]bool {
	result := make(map[string]bool)
	for i := 0; i < len(paths); i++ {
		result[paths[i][index]] = true
	}

	return result
}

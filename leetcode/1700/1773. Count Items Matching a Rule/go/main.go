package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	key, result := keyToInt(ruleKey), 0

	for i := 0; i < len(items); i++ {
		if items[i][key] == ruleValue {
			result++
		}
	}

	return result
}

func keyToInt(key string) int {
	switch key {
	case "type":
		return 0
	case "color":
		return 1
	default:
		return 2
	}
}

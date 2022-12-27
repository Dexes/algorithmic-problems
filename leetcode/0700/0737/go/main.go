package main

func areSentencesSimilarTwo(sentence1 []string, sentence2 []string, pairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}

	pairs, IDs := makeIDs(pairs)
	sets := toSets(pairs, IDs)

	for i := 0; i < len(sentence1); i++ {
		firstID, secondID := IDs[sentence1[i]], IDs[sentence2[i]]
		if firstID == 0 && secondID == 0 {
			if sentence1[i] != sentence2[i] {
				return false
			}

			continue
		}

		if find(firstID, sets) != find(secondID, sets) {
			return false
		}
	}

	return true
}

func find(id int, sets []int) int {
	for sets[id] != sets[sets[id]] {
		sets[id] = sets[sets[id]]
	}

	return sets[id]
}

func makeIDs(pairs [][]string) ([][]string, map[string]int) {
	wIndex, words := 0, make(map[string]int)
	for _, pair := range pairs {
		if pair[0] == pair[1] {
			continue
		}

		if _, ok := words[pair[0]]; !ok {
			words[pair[0]] = len(words) + 1
		}

		if _, ok := words[pair[1]]; !ok {
			words[pair[1]] = len(words) + 1
		}

		pairs[wIndex], wIndex = pair, wIndex+1
	}

	return pairs[:wIndex], words
}

func toSets(pairs [][]string, IDs map[string]int) []int {
	result := make([]int, len(IDs)+1)
	for _, pair := range pairs {
		firstID, secondID := IDs[pair[0]], IDs[pair[1]]
		switch firstParent, secondParent := result[firstID], result[secondID]; {
		case firstParent == 0 && secondParent == 0:
			result[firstID], result[secondID] = firstID, firstID
		case firstParent == 0:
			result[firstID] = secondParent
		case secondParent == 0:
			result[secondID] = firstParent
		default:
			if firstSet, secondSet := find(firstID, result), find(secondID, result); firstSet != secondSet {
				result[firstSet] = secondSet
			}
		}
	}

	return result
}

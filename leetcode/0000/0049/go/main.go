package main

type mask [26]byte

func groupAnagrams(strings []string) [][]string {
	anagrams := make(map[mask][]string)
	for i := 0; i < len(strings); i++ {
		insert(anagrams, strings[i])
	}

	return toArray(anagrams)
}

func insert(anagrams map[mask][]string, s string) {
	m := toMask(s)
	list := anagrams[m]

	if list == nil {
		list = make([]string, 0, 32)
	}

	anagrams[m] = append(list, s)
}

func toMask(s string) (result mask) {
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}

	return result
}

func toArray(anagrams map[mask][]string) [][]string {
	result := make([][]string, 0, len(anagrams))
	for _, x := range anagrams {
		result = append(result, x)
	}

	return result
}

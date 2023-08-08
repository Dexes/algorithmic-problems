package main

type mask [26]byte

func groupAnagrams(strings []string) [][]string {
	anagrams := make(map[mask][]string)
	for _, s := range strings {
		insert(anagrams, s)
	}

	result := make([][]string, 0, len(anagrams))
	for _, x := range anagrams {
		result = append(result, x)
	}

	return result
}

func insert(anagrams map[mask][]string, s string) {
	var m mask
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}

	anagrams[m] = append(anagrams[m], s)
}

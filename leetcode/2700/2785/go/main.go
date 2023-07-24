package main

const capacity = 'u' + 1

var (
	vowels  = []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	isVowel = []bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'z': false}
)

func sortVowels(s string) string {
	var (
		letters   [capacity]int
		prevVowel byte
		isSorted  = true
	)

	for i := 0; i < len(s); i++ {
		if !isVowel[s[i]] {
			continue
		}

		if s[i] < prevVowel {
			isSorted = false
		}

		letters[s[i]]++
		prevVowel = s[i]
	}

	if isSorted {
		return s
	}

	t, vowelIndex := []byte(s), 0
	for i := 0; i < len(t); i++ {
		if !isVowel[t[i]] {
			continue
		}

		for ; letters[vowels[vowelIndex]] == 0; vowelIndex++ {
		}

		letters[vowels[vowelIndex]]--
		t[i] = vowels[vowelIndex]
	}

	return string(t)
}

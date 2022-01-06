package main

var morse = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
	"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	set := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		key := make([]byte, 0, 48)
		for j := 0; j < len(words[i]); j++ {
			key = append(key, morse[words[i][j]-'a']...)
		}

		set[string(key)] = true
	}

	return len(set)
}

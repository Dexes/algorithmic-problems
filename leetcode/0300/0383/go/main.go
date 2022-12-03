package main

func canConstruct(ransomNote string, magazine string) bool {
	counter := make([]int, 123)
	for i := 0; i < len(magazine); i++ {
		counter[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if counter[ransomNote[i]] == 0 {
			return false
		}

		counter[ransomNote[i]]--
	}

	return true
}

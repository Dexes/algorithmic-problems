package main

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence1) > len(sentence2) {
		sentence1, sentence2 = sentence2, sentence1
	}

	left, right1, right2 := -1, len(sentence1), len(sentence2)
	for i := 0; ; i++ {
		if i == len(sentence1) {
			if i == len(sentence2) || sentence2[i] == ' ' {
				return true
			}

			break
		}

		if sentence1[i] != sentence2[i] {
			break
		}

		if sentence1[i] == ' ' {
			left = i
		}
	}

	for i, j := len(sentence1)-1, len(sentence2)-1; ; i, j = i-1, j-1 {
		if i < 0 {
			return j < 0 || sentence2[j] == ' '
		}

		if sentence1[i] != sentence2[j] {
			break
		}

		if sentence1[i] == ' ' {
			right1, right2 = i, j
		}
	}

	return left >= right1 || left >= right2
}

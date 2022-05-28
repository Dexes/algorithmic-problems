package main

func largestWordCount(messages []string, senders []string) string {
	counter := make(map[string]int)
	for i := 0; i < len(messages); i++ {
		counter[senders[i]] += countWords(messages[i])
	}

	result, resultCount := senders[0], counter[senders[0]]
	for i := 0; i < len(senders); i++ {
		if x := counter[senders[i]]; x > resultCount || x == resultCount && senders[i] > result {
			result, resultCount = senders[i], x
		}
	}

	return result
}

func countWords(s string) (result int) {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result++
		}
	}

	return result + 1
}

package main

func replaceWords(dictionary []string, sentence string) string {
	result, trie := make([]byte, 0, len(sentence)+1), makeTrie(dictionary)

	for i := 0; i < len(sentence); i++ {
		word := getWord(sentence, trie, i)
		result = append(result, word...)
		result = append(result, ' ')

		for ; i < len(sentence) && sentence[i] != ' '; i++ {
		}
	}

	return string(result[:len(result)-1])
}

func getWord(sentence string, trie *TrieNode, startIndex int) string {
	i := startIndex
	for ; i < len(sentence) && sentence[i] != ' '; i++ {
		switch nextIndex := sentence[i] - 'a'; {
		case trie.next[nextIndex] == nil:
			for ; i < len(sentence) && sentence[i] != ' '; i++ {
			}

			return sentence[startIndex:i]

		case trie.next[nextIndex].last:
			return sentence[startIndex : i+1]

		default:
			trie = trie.next[nextIndex]
		}
	}

	return sentence[startIndex:i]
}

func makeTrie(words []string) *TrieNode {
	result := &TrieNode{}
	for i := 0; i < len(words); i++ {
		trieInsert(words[i], result)
	}

	return result
}

func trieInsert(s string, node *TrieNode) {
	var next *TrieNode
	for i := 0; i < len(s); i++ {
		if next = node.next[s[i]-'a']; next == nil {
			next = &TrieNode{}
			node.next[s[i]-'a'] = next
		}

		node = next
	}

	node.last = true
}

const abcSize = 26

type TrieNode struct {
	last bool
	next [abcSize]*TrieNode
}

package main

func indexPairs(text string, words []string) [][]int {
	result, trie := make([][]int, 0, len(words)), makeTrie(words)
	for i := 0; i < len(text); i++ {
		result = find(text, i, trie, result)
	}

	return result
}

func find(text string, startIndex int, trie *TrieNode, result [][]int) [][]int {
	for i := startIndex; i < len(text); i++ {
		if trie = trie.next[text[i]-'a']; trie == nil {
			return result
		}

		if trie.last {
			result = append(result, []int{startIndex, i})
		}
	}

	return result
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

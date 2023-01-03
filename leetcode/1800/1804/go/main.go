package main

const abcSize = 26

type TrieNode struct {
	lastNumber int
	number     int
	next       [abcSize]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(word string) {
	var current, next *TrieNode = t.root, nil

	for i := 0; i < len(word); i++ {
		if next = current.next[word[i]-'a']; next == nil {
			next = &TrieNode{}
			current.next[word[i]-'a'] = next
		}

		next.number++
		current = next
	}

	current.lastNumber++
}

func (t *Trie) CountWordsEqualTo(word string) int {
	if x := t.find(word); x != nil {
		return x.lastNumber
	}

	return 0
}

func (t *Trie) CountWordsStartingWith(prefix string) int {
	if x := t.find(prefix); x != nil {
		return x.number
	}

	return 0
}

func (t *Trie) Erase(word string) {
	var current, next *TrieNode = t.root, nil

	for i := 0; i < len(word); i++ {
		if next = current.next[word[i]-'a']; next.number == 1 {
			current.next[word[i]-'a'] = nil
			return
		}

		next.number--
		current = next
	}

	current.lastNumber--
}

func (t *Trie) find(s string) *TrieNode {
	var current, next *TrieNode = t.root, nil

	for i := 0; i < len(s); i++ {
		if next = current.next[s[i]-'a']; next == nil {
			return nil
		}

		current = next
	}

	return current
}

package main

const abcSize = 26

type TrieNode struct {
	last bool
	next [abcSize]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(word string) {
	var (
		current = t.root
		next    *TrieNode
	)

	for i := 0; i < len(word); i++ {
		if next = current.next[word[i]-'a']; next == nil {
			next = &TrieNode{}
			current.next[word[i]-'a'] = next
		}

		current = next
	}

	current.last = true
}

func (t *Trie) Search(word string) bool {
	x := t.find(word)
	return x != nil && x.last
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}

func (t *Trie) find(s string) *TrieNode {
	var (
		current = t.root
		next    *TrieNode
	)

	for i := 0; i < len(s); i++ {
		if next = current.next[s[i]-'a']; next == nil {
			return nil
		}

		current = next
	}

	return current
}

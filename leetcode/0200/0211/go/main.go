package main

const abcSize = 26

type DictionaryNode struct {
	last bool
	next [abcSize]*DictionaryNode
}

type WordDictionary struct {
	root *DictionaryNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &DictionaryNode{}}
}

func (d *WordDictionary) AddWord(word string) {
	var (
		current = d.root
		next    *DictionaryNode
	)

	for i := 0; i < len(word); i++ {
		if next = current.next[word[i]-'a']; next == nil {
			next = &DictionaryNode{}
			current.next[word[i]-'a'] = next
		}

		current = next
	}

	current.last = true
}

func (d *WordDictionary) Search(word string) bool {
	return d.search(d.root, word, 0)
}

func (d *WordDictionary) search(parent *DictionaryNode, word string, index int) bool {
	switch anyLetter, lastIndex, children := word[index] == '.', index == len(word)-1, parent.next; {
	case anyLetter && lastIndex:
		for i := 0; i < abcSize; i++ {
			if children[i] != nil && children[i].last {
				return true
			}
		}

		return false

	case anyLetter:
		for i := 0; i < abcSize; i++ {
			if children[i] != nil && d.search(children[i], word, index+1) {
				return true
			}
		}

		return false

	case lastIndex:
		child := children[word[index]-'a']
		return child != nil && child.last

	default:
		child := children[word[index]-'a']
		return child != nil && d.search(child, word, index+1)
	}
}

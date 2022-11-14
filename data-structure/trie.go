package main

// when to use:
//	- word search
// example: https://leetcode.com/problems/implement-trie-prefix-tree

type Trie struct {
	Root *Node
}

type Node struct {
	IsWord   bool
	Children map[rune]*Node
}

func NewTrie() Trie {
	return Trie{
		Root: &Node{
			Children: make(map[rune]*Node),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, c := range word {
		if _, ok := node.Children[c]; !ok {
			node.Children[c] = &Node{Children: make(map[rune]*Node)}
		}

		node = node.Children[c]
	}

	node.IsWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, c := range word {
		if _, ok := node.Children[c]; !ok {
			return false
		}

		node = node.Children[c]
	}

	if node.IsWord {
		return true
	}

	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, c := range prefix {
		if _, ok := node.Children[c]; !ok {
			return false
		}

		node = node.Children[c]
	}

	return true
}

func (t *Trie) GetWords(prefix string) []string {
	var words []string
	node := t.Root

	// find the node with specified prefix
	for _, c := range prefix {
		if node.Children == nil {
			break
		}

		if _, ok := node.Children[c]; !ok {
			return []string{}
		}

		node = node.Children[c]
	}

	GetWordsUnderNode(node, prefix, &words)

	return words
}

func GetWordsUnderNode(node *Node, curWord string, words *[]string) {
	if node == nil {
		return
	}

	if node.IsWord {
		*words = append(*words, curWord)
	}

	for c, n := range node.Children {
		GetWordsUnderNode(n, curWord+string(c), words)
	}
}

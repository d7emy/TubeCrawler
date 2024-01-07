package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node != nil && node.isEnd
}

func test() {
	words := []string{"apple", "banana", "orange", "grape", "pineapple", "pear"}

	trie := NewTrie()

	for _, word := range words {
		trie.Insert(word)
	}

	wordToCheck := "pear"
	if trie.Search(wordToCheck) {
		fmt.Printf("%s exists in the list of words\n", wordToCheck)
	} else {
		fmt.Printf("%s does not exist in the list of words\n", wordToCheck)
	}
}

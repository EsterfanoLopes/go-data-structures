package main

import "fmt"

const (
	// AlphabetSize is the number of posible characters in the trie
	AlphabetSize = 26
)

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create a new Trie
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	var (
		currentNode = t.root
	)

	for _, l := range w {
		// charIndex is the representation number of the given character starting from index 1
		charIndex := l - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true if the word is included in the trie
func (t *Trie) Search(w string) bool {
	var (
		currentNode = t.root
	)

	for _, l := range w {
		charIndex := l - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()
	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"oregon",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	// Only accept non capital non special letters
	fmt.Println(myTrie.Search(""))
}

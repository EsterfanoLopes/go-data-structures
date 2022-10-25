package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int) {
	if k > n.Key {
		// move right
		if n.Right == nil { // empty spot
			n.Right = &Node{Key: k}
		} else { // already exists
			n.Right.Insert(k)
		}
	} else if k < n.Key {
		// move left
		if n.Left == nil { // empty spot
			n.Left = &Node{Key: k}
		} else { // already exists
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value and RETURN true if there is a node with that value with the depth in the tree
func (n *Node) Search(k int) bool {
	count++
	if n == nil { // not found
		return false
	}

	if k > n.Key {
		// move right
		return n.Right.Search(k)
	} else if k < n.Key {
		// move left
		return n.Left.Search(k)
	} else {
		return true // found
	}
}

func main() {
	tree := &Node{Key: 100}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)

	tree.Insert(32)
	tree.Insert(48)
	tree.Insert(47)

	// exist
	fmt.Println(tree.Search(32), count)
	count = 0

	// do not exist
	fmt.Println(tree.Search(132), count)
}

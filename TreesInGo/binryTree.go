package main

import "fmt"

//node

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search will take a key value
// and reeturn truee if theere is a node with that value

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}

	tree.Insert(50)
	tree.Insert(540)
	tree.Insert(30)
	tree.Insert(770)
	tree.Insert(430)
	tree.Insert(210)
	tree.Insert(20)
	tree.Insert(880)
	tree.Insert(50)
	tree.Insert(2)

	fmt.Println(tree.Search(770))
	fmt.Println(tree.Search(333))
}

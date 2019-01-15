package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(v int) error {
	if n == nil {
		return errors.New("can not insert")
	}
	switch {
	case v == n.value:
		return nil
	case v < n.value:
		if n.left == nil {
			n.left = &Node{v, nil, nil}
			return nil
		}
		return n.left.insert(v)
	case v > n.value:
		if n.right == nil {
			n.right = &Node{v, nil, nil}
			return nil
		}
		return n.right.insert(v)
	}
	return nil
}

func (n *Node) Find(v int) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	switch {
	case v == n.value:
		return n, true

	case v < n.value:
		return n.left.Find(v)

	default:
		return n.right.Find(v)
	}
}

func (n *Node) PrintTree() {
	if n.left != nil {
		n.left.PrintTree()
	}
	fmt.Println(n.value)
	if n.right != nil {
		n.right.PrintTree()
	}
}

func main() {
	tree := Node{46, nil, nil}
	tree.insert(16)
	tree.insert(50)
	tree.insert(40)
	tree.insert(7)
	tree.insert(12)
	tree.insert(15)
	tree.insert(10)
	tree.insert(70)

	tree.PrintTree()

}

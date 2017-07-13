package main

import "fmt"

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func buildTree() Tree {
	//return &Tree{4, Tree{2, Tree{1, nil, nil}, Tree{3, nil, nil}}, Tree{7, Tree{6, nil, nil}, Tree{9, nil, nil}}}
	t := Tree{3, nil, nil}
	t1 := Tree{5, nil, nil}
	t.left = &t1
	return t
}

func traverseTree(t Tree) {
	fmt.Println(t.value)
	if t.left != nil {
		traverseTree(*t.left)
	}
	if t.right != nil {
		traverseTree(*t.right)
	}
}

func main() {
	t := buildTree()
	//fmt.Println(t)
	//fmt.Println(*t.left)
	traverseTree(t)
}

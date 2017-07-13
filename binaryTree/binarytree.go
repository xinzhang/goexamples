package main

import "fmt"

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func buildTree() Tree {
	//return &Tree{4, Tree{2, Tree{1, nil, nil}, Tree{3, nil, nil}}, Tree{7, Tree{6, nil, nil}, Tree{9, nil, nil}}}
	t := Tree{4, nil, nil}
	t1 := Tree{2, nil, nil}
	t2 := Tree{7, nil, nil}
	t.left = &t1
	t.right = &t2

	t3 := Tree{1, nil, nil}
	t4 := Tree{3, nil, nil}
	t1.left = &t3
	t1.right = &t4

	t5 := Tree{6, nil, nil}
	t6 := Tree{9, nil, nil}
	t2.left = &t5
	t2.right = &t6

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

func reverseTree(t Tree) {
	//fmt.Println(t.value)
	fmt.Println("----------------")
	if t.left != nil && t.right != nil {
		tmp := t.left
		t.left = t.right
		t.right = tmp

		fmt.Println(t)
		fmt.Println(t.left)
		fmt.Println(t.right)

		reverseTree(*t.left)
		reverseTree(*t.right)
	}
}

func main() {
	t := buildTree()
	fmt.Println(t)

	//fmt.Println(t)
	//fmt.Println(*t.left)
	reverseTree(t)

	fmt.Println(t)
}

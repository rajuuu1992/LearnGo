package main

import "fmt"

type Complex struct {
	Number int
	Text   string
}

type Tree struct {
	left, right *Tree
	Leaf
	b Branch
}

type Leaf struct {
	val int
}

type Branch struct {
	val string
}

func main() {
	s := Tree{nil, nil, Leaf{8}, Branch{"99aab"}}

	s.val = 10
	s.b.val = "new val"
	fmt.Println(s)

	vals := []int{1, 100, 10, 15, 42, 33, 123, 14, 13}

	Sort(vals)
}

func add(t *Tree, val int) *Tree {
	if t == nil {
		t = new(Tree)
		t.val = val
		return t
	}
	if val < t.val {
		t.left = add(t.left, val)
	} else {
		t.right = add(t.right, val)
	}
	return t
}

func Sort(values []int) {
	var root *Tree

	for _, val := range values {
		root = add(root, val)
	}
	Inorder(root)
}

func Inorder(tree *Tree) {
	if tree == nil {
		return
	}

	Inorder(tree.left)
	fmt.Println(tree.val)
	Inorder(tree.right)
}

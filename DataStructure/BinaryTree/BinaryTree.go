package BinaryTree

import "fmt"

type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	IsRoot bool
}

func PreOrder(t *Tree) {
	if t == nil {
		return
	}

	fmt.Println(t.Value)
	PreOrder(t.Left)
	PreOrder(t.Right)
}

func InOrder(t *Tree) {
	if t == nil {
		return
	}

	InOrder(t.Left)
	fmt.Println(t.Value)
	InOrder(t.Right)
}

func PostOrder(t *Tree) {
	if t == nil {
		return
	}

	PostOrder(t.Left)
	PostOrder(t.Right)
	fmt.Println(t.Value)
}

package BinaryTree

import (
	"fmt"
	"github.com/zbml/DataStructureAndAlgorithm/DataStructure/Stack"
)

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

func PreOrder2(t *Tree) {
	cur := t
	stack := Stack.NewStack()
	for {
		fmt.Println(cur.Value)
		if cur.Left != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			if stack.IsEmpty() {
				break
			}
			cTree, _ := stack.Pop()
			cur = cTree.(*Tree).Right
		}
	}
}

func InOrder(t *Tree) {
	if t == nil {
		return
	}

	InOrder(t.Left)
	fmt.Println(t.Value)
	InOrder(t.Right)
}

func InOrder2(t *Tree) {
	cur := t
	s := Stack.NewStack()
	for {
		if cur.Left != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			fmt.Println(cur.Value)
			if s.IsEmpty() {
				break
			}
			t, _ := s.Pop()
			cTree := t.(*Tree)
			fmt.Println(cTree.Value)
			cur = cTree.Right
		}
	}
}

func PostOrder(t *Tree) {
	if t == nil {
		return
	}

	PostOrder(t.Left)
	PostOrder(t.Right)
	fmt.Println(t.Value)
}

func PostOrder2(t *Tree) {
	cur := t
	s := Stack.NewStack()
	lastVisit := t

	for cur != nil || !s.IsEmpty() {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		node, _ := s.Peek()
		cur = node.(*Tree)
		if cur.Right == nil || cur.Right == lastVisit {
			fmt.Println(cur.Value)
			s.Pop()
			lastVisit = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
}

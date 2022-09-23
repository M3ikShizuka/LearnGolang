package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *Tree, ch chan int) {
	// Inorder Traversal
	if t == nil {
		return
	}

	// Check left tree
	Walk(t.Left, ch)

	// Send to channel
	ch <- t.Value

	// Check right tree
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	t1Chan, t2Chan := make(chan int), make(chan int)

	go Walk(t1, t1Chan)
	go Walk(t2, t2Chan)

	for {
		t1Value, t1ChanOk := <-t1Chan
		t2Value, t2ChanOk := <-t2Chan

		if t1Value != t2Value ||
			t1ChanOk != t2ChanOk {
			return false
		}

		if t1ChanOk == false && t2ChanOk == false {
			break
		}
	}

	return true
}

func main() {
	tree1 := New(1)
	ch := make(chan int)

	go Walk(tree1, ch)

	tree2Like1 := New(1)
	fmt.Println(Same(tree1, tree2Like1))

	go Walk(tree2Like1, ch)

	tree2Diff1 := New(2)
	fmt.Println(Same(tree1, tree2Diff1))

	go Walk(tree2Diff1, ch)

	for value := range ch {
		fmt.Println(value)
	}
}

package avl

import (
	"fmt"
)

type Key struct {
	val int
}

func (a *Key) Compare(b Interface) int {
	bval := b.(*Key).val
	switch {
	case a.val > bval:
		return 1
	case a.val < bval:
		return -1
	}
	return 0
}

func parentKey(node *Node) int {
	if node.Parent != nil {
		return node.Parent.Key.(*Key).val
	}
	return -1
}

func biggerKey(node *Node) int {
	if node.Bigger != nil {
		return node.Bigger.Key.(*Key).val
	}
	return -1
}

func thisKey(node *Node) int {
	if node != nil {
		return node.Key.(*Key).val
	}
	return -1
}

func Test() {
	tree := Create()
	keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
	for _, key := range keys {
		Insert(tree, &Key{key})
	}

	Walk(tree, func(node *Node) {
		fmt.Println(thisKey(node), parentKey(node), Height(node),
		biggerKey(node))
	})

	node := FindMinimum(*tree)
	for node != nil {
		fmt.Println(thisKey(node))
		node = node.Bigger
	}
}

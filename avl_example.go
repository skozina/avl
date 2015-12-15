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

func Example() {
	tree := Create()
	keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
	for _, key := range keys {
		fmt.Println("Insert ", key)
		Insert(tree, &Key{key})
	}

	fmt.Println("Walk the tree:")
	Walk(tree, func(node *Node) {
		fmt.Println("Node: ", thisKey(node),
		    "Parent: ", parentKey(node),
		    "Height: ", Height(node),
		    "Next bigger: ", biggerKey(node))
	})

	fmt.Println("Walk tree from smallest node:")
	node := FindMinimum(*tree)
	for node != nil {
		fmt.Println("Node: ", thisKey(node))
		node = node.Bigger
	}
}

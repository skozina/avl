package avl

import (
	"fmt"
)

type avlNode struct {
	Key			Interface
	Height			int
	Lchild, Rchild, Parent	*avlNode
	Smaller, Bigger		*avlNode // Neighbor in sorted array
}

type Interface interface {
	Compare(b Interface) int
}

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

func leftRotate(root *avlNode) *avlNode {
	node := root.Rchild

	root.Rchild = node.Lchild
	if root.Rchild != nil {
		root.Rchild.Parent = root
	}

	node.Lchild = root
	node.Parent = root.Parent
	root.Parent = node

	root.Height = max(Height(root.Lchild), Height(root.Rchild)) + 1
	node.Height = max(Height(node.Rchild), Height(node.Lchild)) + 1
	return node
}

func leftRigthRotate(root *avlNode) *avlNode {
	root.Lchild = leftRotate(root.Lchild)
	root = rightRotate(root)
	return  root
}

func rightRotate(root *avlNode) *avlNode {
	node := root.Lchild

	root.Lchild = node.Rchild
	if root.Lchild != nil {
		root.Lchild.Parent = root
	}

	node.Rchild = root
	node.Parent = root.Parent
	root.Parent = node

	root.Height = max(Height(root.Lchild), Height(root.Rchild)) + 1
	node.Height = max(Height(node.Lchild), Height(node.Rchild)) + 1
	return node
}

func rightLeftRotate(root *avlNode) *avlNode {
	root.Rchild = rightRotate(root.Rchild)
	root = leftRotate(root)
	return  root
}

func Height(root *avlNode) int {
	if root != nil {
		return root.Height
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * TODO tohle se musi volat jenom jednou pro kazdy node, ne po kazdem vyvazeni v
 * rekurzi!
 */
func updateLinks(node *avlNode) {
	/* Fix the links to sorted neighbors */
	neighbor := findSuccessor(node)
	if neighbor != nil {
		node.Bigger = neighbor
		if neighbor.Smaller != nil {
			node.Smaller = neighbor.Smaller
			neighbor.Smaller.Bigger = node
		}
		neighbor.Smaller = node
	} else {
		neighbor = findPredecessor(node)
		if neighbor != nil {
			node.Smaller = neighbor
			if neighbor.Bigger != nil {
				node.Bigger = neighbor.Bigger
				neighbor.Bigger.Smaller = node
			}
			neighbor.Bigger = node
		}
	}

}

func Insert(root *avlNode, key Interface) *avlNode {
	if root == nil {
		root = &avlNode{key, 0, nil, nil, nil, nil, nil}
		root.Height = max(Height(root.Lchild), Height(root.Rchild)) + 1
		return root
	}

	if key.Compare(root.Key) < 0 {
		root.Lchild = Insert(root.Lchild, key)
		root.Lchild.Parent = root
		updateLinks(root.Lchild)
		if Height(root.Lchild)-Height(root.Rchild) == 2 {
			if key.Compare(root.Lchild.Key) < 0 {
				root = rightRotate(root)
			} else {
				root = leftRigthRotate(root)
			}
		}
	} 

	if key.Compare(root.Key) > 0 {
		root.Rchild = Insert(root.Rchild, key)
		root.Rchild.Parent = root
		updateLinks(root.Rchild)
		if Height(root.Rchild)-Height(root.Lchild) == 2 {
			if key.Compare(root.Rchild.Key) > 0 {
				root = leftRotate(root)
			} else {
				root = rightLeftRotate(root)
			}
		}
	}

	root.Height = max(Height(root.Lchild), Height(root.Rchild)) + 1
	return root
}

func findMinimum(node *avlNode) *avlNode {
	if node.Lchild != nil {
		return findMinimum(node.Lchild)
	}
	return node
}

func findMaximum(node *avlNode) *avlNode {
	if node.Rchild != nil {
		return findMaximum(node.Rchild)
	}
	return node
}

/*
 * Find node with value next bigger to the current node.
 */
func findSuccessor(node *avlNode) *avlNode {
	if node.Rchild != nil {
		fmt.Println("successor ", thisKey(node),
		    " rchild != nil")
		return findMinimum(node.Rchild)
	}

	parent := node.Parent
	this := node
	for parent != nil && this == parent.Rchild {
		this = parent
		parent = this.Parent
	}

	fmt.Println("successor ", thisKey(node), " is ", thisKey(parent))
	return parent
}

/*
 * Find node with value next bigger to the current node.
 */
func findPredecessor(node *avlNode) *avlNode {
	if node.Lchild != nil {
		fmt.Println("predecessor ", thisKey(node),
		    " lchild != nil")
		return findMaximum(node.Lchild)
	}

	parent := node.Parent
	this := node
	for parent != nil && this == parent.Lchild {
		this = parent
		parent = this.Parent
	}

	fmt.Println("predecessor ", thisKey(node), " is ", thisKey(parent))
	return parent
}

type action func(node *avlNode)

func inOrder(root *avlNode, action action) {
	if root == nil {
		return
	}

	inOrder(root.Lchild, action)
	action(root)
	inOrder(root.Rchild, action)
}

func parentKey(node *avlNode) int {
	if node.Parent != nil {
		return node.Parent.Key.(*Key).val
	}
	return -1
}

func biggerKey(node *avlNode) int {
	if node.Bigger != nil {
		return node.Bigger.Key.(*Key).val
	}
	return -1
}

func thisKey(node *avlNode) int {
	if node != nil {
		return node.Key.(*Key).val
	}
	return -1
}

func AvlTest() {
	var root *avlNode
 //	keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
 	keys := []int{2, 6, 5}
	for _, key := range keys {
		root = Insert(root, &Key{key})
	}

	inOrder(root, func(node *avlNode) {
		fmt.Println(thisKey(node), parentKey(node), Height(node),
		biggerKey(node))
	})

	node := findMinimum(root)
	for node != nil {
		fmt.Println(thisKey(node))
		if thisKey(node) == 6 {
			panic(1)
		}
		node = node.Bigger
	}
}

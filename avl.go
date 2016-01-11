package avl

/*
 * Simple AVL tree with parent pointers in node included.
 * All nodes are also put on a sorted linked list. This allows search for range
 * of values.
 */

type Node struct {
	Key                    Interface
	Height                 int
	Lchild, Rchild, Parent *Node
	Smaller, Bigger        *Node // Neighbor in sorted array
}

type Tree **Node

type Interface interface {
	Compare(b Interface) int
}

func leftRotate(root *Node) *Node {
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

func leftRigthRotate(root *Node) *Node {
	root.Lchild = leftRotate(root.Lchild)
	root = rightRotate(root)
	return root
}

func rightRotate(root *Node) *Node {
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

func rightLeftRotate(root *Node) *Node {
	root.Rchild = rightRotate(root.Rchild)
	root = leftRotate(root)
	return root
}

func Height(root *Node) int {
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

func updateLinks(node *Node) {
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

func insertTree(tree **Node, key Interface) *Node {
	root := *tree
	var node *Node

	if root == nil {
		root = &Node{key, 0, nil, nil, nil, nil, nil}
		root.Height = max(Height(root.Lchild), Height(root.Rchild)) + 1
		*tree = root
		return root
	}

	if key.Compare(root.Key) < 0 {
		node = insertTree(&root.Lchild, key)
		root.Lchild.Parent = root
		if Height(root.Lchild)-Height(root.Rchild) == 2 {
			if key.Compare(root.Lchild.Key) < 0 {
				root = rightRotate(root)
			} else {
				root = leftRigthRotate(root)
			}
		}
	}

	if key.Compare(root.Key) > 0 {
		node = insertTree(&root.Rchild, key)
		root.Rchild.Parent = root
		if Height(root.Rchild)-Height(root.Lchild) == 2 {
			if key.Compare(root.Rchild.Key) > 0 {
				root = leftRotate(root)
			} else {
				root = rightLeftRotate(root)
			}
		}
	}

	root.Height = max(Height(root.Lchild), Height(root.Rchild)) + 1
	*tree = root
	return node
}

func Insert(tree Tree, key Interface) *Node {
	node := insertTree(tree, key)
	if node != nil {
		updateLinks(node)
	}
	return node
}

/*
 * The left-most node from the current node.
 */
func FindMinimum(node *Node) *Node {
	if node.Lchild != nil {
		return FindMinimum(node.Lchild)
	}
	return node
}

/*
 * The right-most node from the current node.
 */
func FindMaximum(node *Node) *Node {
	if node.Rchild != nil {
		return FindMaximum(node.Rchild)
	}
	return node
}

/*
 * Find node with value next bigger to the current node.
 */
func findSuccessor(node *Node) *Node {
	if node.Rchild != nil {
		return FindMinimum(node.Rchild)
	}

	parent := node.Parent
	this := node
	for parent != nil && this == parent.Rchild {
		this = parent
		parent = this.Parent
	}

	return parent
}

/*
 * Find node with value next bigger to the current node.
 */
func findPredecessor(node *Node) *Node {
	if node.Lchild != nil {
		return FindMaximum(node.Lchild)
	}

	parent := node.Parent
	this := node
	for parent != nil && this == parent.Lchild {
		this = parent
		parent = this.Parent
	}

	return parent
}

type Walker func(node *Node) bool

/*
 * Call given function on all nodes in the tree.
 * Depth-first walk.
 */
func Walk(tree Tree, action Walker) bool {
	root := *tree
	if root == nil {
		return true
	}

	if Walk(&root.Lchild, action) == false {
		return false
	}
	if action(root) == false {
		return false
	}
	if Walk(&root.Rchild, action) == false {
		return false
	}
	return true
}

/*
 * Empty tree is nil.
 */
func Create() Tree {
	var result *Node
	return &result
}

/*
 * Find given node or the one next smaller.
 */
func FindSmaller(tree Tree, key Interface) *Node {
	root := *tree

	if root == nil {
		return root
	}

	if key.Compare(root.Key) < 0 {
		if root.Lchild != nil {
			return FindSmaller(&root.Lchild, key)
		}
		return root.Smaller
	}
	if key.Compare(root.Key) > 0 {
		if root.Rchild != nil {
			return FindSmaller(&root.Rchild, key)
		}
		return root
	}
	return root
}

/*
 * Find given node or the one next bigger.
 */
func FindBigger(tree Tree, key Interface) *Node {
	root := *tree

	if root == nil {
		return root
	}

	if key.Compare(root.Key) < 0 {
		if root.Lchild != nil {
			return FindBigger(&root.Lchild, key)
		}
		return root
	}
	if key.Compare(root.Key) > 0 {
		if root.Rchild != nil {
			return FindBigger(&root.Rchild, key)
		}
		return root.Bigger
	}
	return root
}

package tree

// Tree represents a tree.
type Tree struct {
	root *Node
}

// NewTree returns a new Tree object.
func NewTree(root *Node) *Tree {
	return &Tree{
		root: root,
	}
}

// Empty returns true if the tree contains any node. Otherwise false.
func (t *Tree) Empty() bool {
	return t.Size() == 0
}

// Size returns the number of the nodes in the tree.
func (t *Tree) Size() int {
	if t.root == nil {
		return 0
	}

	return t.root.size()
}

// Height returns the number of step to travel from the root of the tree to the furthest leaf.
func (t *Tree) Height() int {
	if t.root == nil {
		return 0
	}

	return t.root.height()
}

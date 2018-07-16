package tree

// Node represents a node in a tree.
type Node struct {
	LeftChild  *Node
	RightChild *Node
	Value      int
}

// NewNode returns a new TreeNode object.
func NewNode(value int) *Node {
	return &Node{
		Value:      value,
		LeftChild:  nil,
		RightChild: nil,
	}
}

// size is a helper function that returns the number of the child nodes of the current node, including the current node.
func (n *Node) size() int {
	size := 1
	if n.LeftChild != nil {
		size += n.LeftChild.size()
	}
	if n.RightChild != nil {
		size += n.RightChild.size()
	}

	return size
}

// height is a helper function that returns the number of steps from the current node to the furthest leaf node.
func (n *Node) height() int {
	leftHeight := 0
	rightHeight := 0
	if n.LeftChild != nil {
		leftHeight = n.LeftChild.height()
	}
	if n.RightChild != nil {
		rightHeight = n.RightChild.height()
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

type serializedNode struct {
	Value int `json:"value"`
	Index int `json:"index"`
}

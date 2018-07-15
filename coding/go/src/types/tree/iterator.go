package tree

import "types/stack"

// Iterator defines the interface for iterating through a tree.
type Iterator interface {
	HasNext() bool
	Next() *Node
}

// PreOrderIterator iterates through a tree in pre-order method.
type PreOrderIterator struct {
	nodeStack stack.Stack
}

// NewPreOrderIterator returns a new PreOrderIterator object.
func NewPreOrderIterator(tree *Tree) Iterator {
	iter := &PreOrderIterator{}

	if !tree.Empty() {
		iter.nodeStack.Push(tree.root)
	}

	return iter
}

// HasNext returns true if there is still un-visited node in the tree. Otherwise false.
func (i *PreOrderIterator) HasNext() bool {
	return !i.nodeStack.Empty()
}

// Next returns the next node to be visited.
func (i *PreOrderIterator) Next() *Node {
	topNode := i.nodeStack.Peek().(*Node)
	i.nodeStack.Pop()

	if topNode.RightChild != nil {
		i.nodeStack.Push(topNode.RightChild)
	}

	if topNode.LeftChild != nil {
		i.nodeStack.Push(topNode.LeftChild)
	}

	return topNode
}

// PreOrderIterator iterates through a tree in in-order method.
type InOrderIterator struct {
	nodeStack stack.Stack
}

// NewInOrderIterator returns a new InOrderIterator object.
func NewInOrderIterator(tree *Tree) Iterator {
	iter := &InOrderIterator{}

	node := tree.root
	for node != nil {
		iter.nodeStack.Push(node)
		node = node.LeftChild
	}

	return iter
}

// HasNext returns true if there is still un-visited node in the tree. Otherwise false.
func (i *InOrderIterator) HasNext() bool {
	return !i.nodeStack.Empty()
}

// Next returns the next node to be visited.
func (i *InOrderIterator) Next() *Node {
	topNode := i.nodeStack.Peek().(*Node)
	i.nodeStack.Pop()

	node := topNode.RightChild
	for node != nil {
		i.nodeStack.Push(node)
		node = node.LeftChild
	}

	return topNode
}

// PostOrderIterator iterates through a tree in post-order method.
type PostOrderIterator struct {
	orderedStack stack.Stack
}

// NewPostOrderIterator returns a new PostOrderIterator object.
func NewPostOrderIterator(tree *Tree) Iterator {
	iter := &PostOrderIterator{}

	if tree.root == nil {
		return iter
	}

	var st stack.Stack
	st.Push(tree.root)
	for !st.Empty() {
		node := st.Peek().(*Node)
		st.Pop()
		iter.orderedStack.Push(node)
		if node.LeftChild != nil {
			st.Push(node.LeftChild)
		}
		if node.RightChild != nil {
			st.Push(node.RightChild)
		}
	}

	return iter
}

// HasNext returns true if there is still un-visited node in the tree. Otherwise false.
func (i *PostOrderIterator) HasNext() bool {
	return !i.orderedStack.Empty()
}

// Next returns the next node to be visited.
func (i *PostOrderIterator) Next() *Node {
	topNode := i.orderedStack.Peek().(*Node)
	i.orderedStack.Pop()
	return topNode
}

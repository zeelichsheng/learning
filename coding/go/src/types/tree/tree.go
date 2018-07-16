package tree

import "encoding/json"

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

// Marshal converts a tree to a JSON string.
func MarshalJSON(t *Tree) ([]byte, error) {
	st := &serializedTree{
		Nodes: t.marshalDfs(t.root, 1),
	}

	return json.Marshal(st)
}

// Unmarshal converts a JSON string to a tree.
func UnmarshalJSON(data []byte) (*Tree, error) {
	var st serializedTree
	err := json.Unmarshal(data, &st)
	if err != nil {
		return nil, err
	}

	var root *Node
	indexMap := make(map[int]*Node)
	for _, serializedNode := range st.Nodes {
		treeNode := &Node{
			Value: serializedNode.Value,
		}

		if parentTreeNode, ok := indexMap[serializedNode.Index/2]; ok {
			if serializedNode.Index%2 == 0 {
				parentTreeNode.LeftChild = treeNode
			} else {
				parentTreeNode.RightChild = treeNode
			}
		}

		indexMap[serializedNode.Index] = treeNode

		if root == nil {
			root = treeNode
		}
	}

	return NewTree(root), nil
}

func (t *Tree) marshalDfs(n *Node, index int) []*serializedNode {
	if n == nil {
		return nil
	}

	sn := &serializedNode{
		Value: n.Value,
		Index: index,
	}

	leftSNs := t.marshalDfs(n.LeftChild, index*2)
	rightSNs := t.marshalDfs(n.RightChild, index*2+1)

	sns := make([]*serializedNode, 0)
	sns = append(sns, sn)
	if leftSNs != nil {
		sns = append(sns, leftSNs...)
	}
	if rightSNs != nil {
		sns = append(sns, rightSNs...)
	}

	return sns
}

type serializedTree struct {
	Nodes []*serializedNode `json:"nodes"`
}

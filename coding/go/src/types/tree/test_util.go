package tree

import "math/rand"

const (
	treeNodeMaxValue        = 100
	singleNodeTreeNodeValue = 1
)

func GetEmptyTree() *Tree {
	return &Tree{
		root: nil,
	}
}

func GetSingleNodeTree() *Tree {
	return &Tree{
		root: NewNode(singleNodeTreeNodeValue),
	}
}

func GetAllLeftNodeTree(height int) *Tree {
	return &Tree{
		root: getOneSideTreeRootNode(height, true),
	}
}

func GetAllRightNodeTree(height int) *Tree {
	return &Tree{
		root: getOneSideTreeRootNode(height, false),
	}
}

func getOneSideTreeRootNode(height int, isLeft bool) *Node {
	var root *Node
	var prevNode *Node

	i := 1
	for i <= height {
		node := &Node{
			Value: rand.Intn(treeNodeMaxValue + 1),
		}
		if root == nil {
			root = node
		}

		if prevNode != nil {
			if isLeft {
				prevNode.LeftChild = node
			} else {
				prevNode.RightChild = node
			}
		}
		prevNode = node
		i++
	}

	return root
}

func GetFullTree(height int) *Tree {
	return &Tree{
		root: getFullTreeRootNode(height),
	}
}

func getFullTreeRootNode(height int) *Node {
	return dfsGetTreeNode(height, func() bool {
		return true
	})
}

func GetRandomTree(height int) *Tree {
	return &Tree{
		root: getRandomTreeRootNode(height),
	}
}

func getRandomTreeRootNode(height int) *Node {
	return dfsGetTreeNode(height, func() bool {
		return rand.Intn(201) <= 100
	})
}

func dfsGetTreeNode(height int, shouldGenerateThisNode func() bool) *Node {
	if height == 0 {
		return nil
	}

	if !shouldGenerateThisNode() {
		return nil
	}
	root := &Node{
		Value: rand.Intn(treeNodeMaxValue + 1),
	}

	root.LeftChild = dfsGetTreeNode(height-1, shouldGenerateThisNode)
	root.RightChild = dfsGetTreeNode(height-1, shouldGenerateThisNode)

	return root

}

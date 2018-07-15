package tree

import (
	"math/rand"
	"testing"
)

const (
	treeNodeMaxValue        = 100
	singleNodeTreeNodeValue = 1
)

func TestTree_HeightOfEmptyTree(t *testing.T) {
	testTreeHeight(t, "empty", 0, func(int) *Tree {
		return GetEmptyTree()
	})
}

func TestTree_HeightOfSingleNodeTree(t *testing.T) {
	testTreeHeight(t, "single node", 1, func(int) *Tree {
		return GetSingleNodeTree()
	})
}

func TestTree_HeightOfAllLeftNodeTree(t *testing.T) {
	testTreeHeight(t, "all left node", 5, GetAllLeftNodeTree)
}

func TestTree_HeightOfAllRightNodeTree(t *testing.T) {
	testTreeHeight(t, "all right node", 3, GetAllRightNodeTree)
}

func TestTree_HeightOfFullTree(t *testing.T) {
	testTreeHeight(t, "full", 4, GetFullTree)
}

func TestTree_HeightOfRandomTree(t *testing.T) {
	expectedHeight := 5
	tree := GetRandomTree(expectedHeight)
	actualHeight := tree.Height()

	if actualHeight > expectedHeight {
		t.Errorf("Testing random tree height. Expect tree height to be equal to or less than %s. Actual %d.", expectedHeight, actualHeight)
	}
}

func testTreeHeight(t *testing.T, treeType string, expectedHeight int, treeGenerator func(int) *Tree) {
	tree := treeGenerator(expectedHeight)
	actualHeight := tree.Height()

	if actualHeight != expectedHeight {
		t.Errorf("Testing %s tree height. Expected %d. Actual %d.", treeType, expectedHeight, actualHeight)
	}
}

func TestTree_SizeOfEmptyTree(t *testing.T) {
	testTreeSize(t, "empty", 0, 0, func(int) *Tree {
		return GetEmptyTree()
	})
}

func TestTree_SizeOfSingleNodeTree(t *testing.T) {
	testTreeSize(t, "single node", 1, 1, func(int) *Tree {
		return GetSingleNodeTree()
	})
}

func TestTree_SizeOfAllLeftNodeTree(t *testing.T) {
	testTreeSize(t, "all left node", 5, 5, GetAllLeftNodeTree)
}

func TestTree_SizeOfAllRightNodeTree(t *testing.T) {
	testTreeSize(t, "all right node", 3, 3, GetAllRightNodeTree)
}

func TestTree_SizeOfFullTree(t *testing.T) {
	height := 4
	expectedSize := 0
	i := height
	for i > 0 {
		expectedSize <<= 1
		expectedSize++
		i--
	}
	testTreeSize(t, "full", height, expectedSize, GetFullTree)
}

func testTreeSize(t *testing.T, treeType string, height int, expectedSize int, treeGenerator func(int) *Tree) {
	tree := treeGenerator(height)
	actualSize := tree.Size()

	if actualSize != expectedSize {
		t.Errorf("Testing %s tree size. Expected %d. Actual %d.", treeType, expectedSize, actualSize)
	}
}

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

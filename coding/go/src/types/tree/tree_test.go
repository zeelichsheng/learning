package tree

import (
	"encoding/json"
	"testing"
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

func TestMarshalJSON_EmptyTree(t *testing.T) {
	testMarshalJSON(t, "empty", GetEmptyTree())
}

func TestMarshalJSON_SingleNodeTree(t *testing.T) {
	testMarshalJSON(t, "single node", GetSingleNodeTree())
}

func TestMarshalJSON_AllLeftNodeTree(t *testing.T) {
	testMarshalJSON(t, "all left node", GetAllLeftNodeTree(5))
}

func TestMarshalJSON_AllRightNodeTree(t *testing.T) {
	testMarshalJSON(t, "all right node", GetAllRightNodeTree(3))
}

func TestMarshalJSON_FullTree(t *testing.T) {
	testMarshalJSON(t, "full", GetFullTree(4))
}

func TestMarshalJSON_RandomTree(t *testing.T) {
	testMarshalJSON(t, "random", GetRandomTree(6))
}

func testMarshalJSON(t *testing.T, treeType string, tr *Tree) {
	ser, err := MarshalJSON(tr)
	if err != nil {
		t.Errorf("Failed to marshal %s tree. Invalid JSON converstion: %+v", treeType, err)
	}

	var str serializedTree
	err = json.Unmarshal(ser, &str)
	if err != nil {
		t.Errorf("Failed to marshal %s tree. Invalid JSON output: %+v", treeType, err)
	}

	if len(str.Nodes) != tr.Size() {
		t.Errorf("Failed to marshal %s tree. Original tree size is %d. Serialized tree size is %d.", treeType, tr.Size(), len(str.Nodes))
	}

	iter := NewPreOrderIterator(tr)
	idx := 0
	for iter.HasNext() {
		n := iter.Next()
		strn := str.Nodes[idx]
		if n.Value != strn.Value {
			t.Errorf("Failed to marshal %s tree. Original tree node value is %d. Serialized tree node value is %d.", treeType, n.Value, strn.Value)
		}
		idx++
	}
}

func TestUnmarshalJSON_EmptyTree(t *testing.T) {
	testUnmarshalJSON(t, "empty", "{}")
}

func TestUnmarshalJSON_SingleNodeTree(t *testing.T) {
	testUnmarshalJSON(t, "single node", `{"nodes":[{"value":1,"index":1}]}`)
}

func TestUnmarshalJSON_AllLeftNodeTree(t *testing.T) {
	testUnmarshalJSON(t, "all left node", `{"nodes":[{"value":65,"index":1},{"value":82,"index":2},{"value":29,"index":4},{"value":87,"index":8},{"value":48,"index":16}]}`)
}

func TestUnmarshalJSON_AllRightNodeTree(t *testing.T) {
	testUnmarshalJSON(t, "all right node", `{"nodes":[{"value":65,"index":1},{"value":82,"index":3},{"value":29,"index":7}]}`)
}

func TestUnmarshalJSON_FullTree(t *testing.T) {
	testUnmarshalJSON(t, "full", `{"nodes":[{"value":65,"index":1},{"value":82,"index":2},{"value":29,"index":4},{"value":87,"index":8},{"value":48,"index":9},{"value":39,"index":5},{"value":37,"index":10},{"value":95,"index":11},{"value":70,"index":3},{"value":48,"index":6},{"value":33,"index":12},{"value":24,"index":13},{"value":38,"index":7},{"value":75,"index":14},{"value":7,"index":15}]}`)
}

func TestUnmarshalJSON_RandomTree(t *testing.T) {
	testUnmarshalJSON(t, "random", `{"nodes":[{"value":82,"index":1},{"value":87,"index":2},{"value":39,"index":4},{"value":70,"index":9}]}`)
}

func testUnmarshalJSON(t *testing.T, treeType string, ser string) {
	tr, err := UnmarshalJSON([]byte(ser))
	if err != nil {
		t.Errorf("Failed to unmarshal %s tree. Invalid JSON conversion: %+v", treeType, err)
	}

	var str serializedTree
	json.Unmarshal([]byte(ser), &str)

	if len(str.Nodes) != tr.Size() {
		t.Errorf("Failed to unmarshal %s tree. Original tree size is %d. Deserialized tree size is %d.", treeType, len(str.Nodes), tr.Size())
	}

	iter := NewPreOrderIterator(tr)
	idx := 0
	for iter.HasNext() {
		n := iter.Next()
		strn := str.Nodes[idx]
		if n.Value != strn.Value {
			t.Errorf("Failed to unmarshal %s tree. Original tree node value is %d. Deserialized tree node value is %d.", treeType, strn.Value, n.Value)
		}
		idx++
	}
}

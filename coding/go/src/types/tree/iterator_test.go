package tree

import "testing"

func TestPreOrderIterator_HasNextOfEmptyTree(t *testing.T) {
	tree := GetEmptyTree()
	iter := NewPreOrderIterator(tree)

	if iter.HasNext() != false {
		t.Errorf("Testing pre-order iterator for empty tree. Wrong indication of hasNext. Expected %t. Actual %t.", false, iter.HasNext())
	}
}

func TestPreOrderIterator_HasNextOfNonEmptyTree(t *testing.T) {
	tree := GetSingleNodeTree()
	iter := NewPreOrderIterator(tree)

	if iter.HasNext() != true {
		t.Errorf("Testing pre-order iterator for empty tree. Wrong indication of hasNext. Expected %t. Actual %t.", true, iter.HasNext())
	}
}

func TestPreOrderIterator_IterSingleTree(t *testing.T) {
	tree := GetSingleNodeTree()
	iter := NewPreOrderIterator(tree)
	examinePreOrderIterator(t, tree.root, iter)
}

func TestPreOrderIterator_IterAllLeftTree(t *testing.T) {
	tree := GetAllLeftNodeTree(5)
	iter := NewPreOrderIterator(tree)
	examinePreOrderIterator(t, tree.root, iter)
}

func TestPreOrderIterator_IterAllRightTree(t *testing.T) {
	tree := GetAllRightNodeTree(4)
	iter := NewPreOrderIterator(tree)
	examinePreOrderIterator(t, tree.root, iter)
}

func TestPreOrderIterator_IterFullTree(t *testing.T) {
	tree := GetFullTree(3)
	iter := NewPreOrderIterator(tree)
	examinePreOrderIterator(t, tree.root, iter)
}

func TestPreOrderIterator_IterRandomTree(t *testing.T) {
	tree := GetRandomTree(6)
	iter := NewPreOrderIterator(tree)
	examinePreOrderIterator(t, tree.root, iter)
}

func examinePreOrderIterator(t *testing.T, node *Node, iter Iterator) {
	examinePreOrderIteratorDfs(t, node, iter)
	if iter.HasNext() {
		t.Error("Iterator should have no more node to iterate.")
	}
}

func examinePreOrderIteratorDfs(t *testing.T, node *Node, iter Iterator) {
	if node == nil {
		return
	}

	if !iter.HasNext() {
		t.Error("Iterator has no next node, but current node is not nil.")
		return
	}
	actualNode := iter.Next()
	if node.Value != actualNode.Value {
		t.Errorf("Iterator returns wrong node. Expected data value is %d. Actual data value is %d.", node.Value, actualNode.Value)
	}

	examinePreOrderIteratorDfs(t, node.LeftChild, iter)
	examinePreOrderIteratorDfs(t, node.RightChild, iter)
}

func TestInOrderIterator_HasNextOfEmptyTree(t *testing.T) {
	tree := GetEmptyTree()
	iter := NewInOrderIterator(tree)

	if iter.HasNext() != false {
		t.Errorf("Testing in-order iterator for empty tree. Wrong indication of hasNext. Expected %t. Actual %t.", false, iter.HasNext())
	}
}

func TestInOrderIterator_HasNextOfNonEmptyTree(t *testing.T) {
	tree := GetSingleNodeTree()
	iter := NewInOrderIterator(tree)

	if iter.HasNext() != true {
		t.Errorf("Testing in-order iterator for empty tree. Wrong indication of hasNext. Expected %t. Actual %t.", true, iter.HasNext())
	}
}

func TestInOrderIterator_IterSingleTree(t *testing.T) {
	tree := GetSingleNodeTree()
	iter := NewInOrderIterator(tree)
	examineInOrderIterator(t, tree.root, iter)
}

func TestInOrderIterator_IterAllLeftTree(t *testing.T) {
	tree := GetAllLeftNodeTree(5)
	iter := NewInOrderIterator(tree)
	examineInOrderIterator(t, tree.root, iter)
}

func TestInOrderIterator_IterAllRightTree(t *testing.T) {
	tree := GetAllRightNodeTree(4)
	iter := NewInOrderIterator(tree)
	examineInOrderIterator(t, tree.root, iter)
}

func TestInOrderIterator_IterFullTree(t *testing.T) {
	tree := GetFullTree(3)
	iter := NewInOrderIterator(tree)
	examineInOrderIterator(t, tree.root, iter)
}

func TestInOrderIterator_IterRandomTree(t *testing.T) {
	tree := GetRandomTree(6)
	iter := NewInOrderIterator(tree)
	examineInOrderIterator(t, tree.root, iter)
}

func examineInOrderIterator(t *testing.T, node *Node, iter Iterator) {
	examineInOrderIteratorDfs(t, node, iter)
	if iter.HasNext() {
		t.Error("Iterator should have no more node to iterate.")
	}
}

func examineInOrderIteratorDfs(t *testing.T, node *Node, iter Iterator) {
	if node == nil {
		return
	}

	examineInOrderIteratorDfs(t, node.LeftChild, iter)

	if !iter.HasNext() {
		t.Error("Iterator has no next node, but current node is not nil.")
		return
	}
	actualNode := iter.Next()
	if node.Value != actualNode.Value {
		t.Errorf("Iterator returns wrong node. Expected data value is %d. Actual data value is %d.", node.Value, actualNode.Value)
	}

	examineInOrderIteratorDfs(t, node.RightChild, iter)
}

func TestPostOrderIterator_HasNextOfEmptyTree(t *testing.T) {
	tree := GetEmptyTree()
	iter := NewPostOrderIterator(tree)

	if iter.HasNext() != false {
		t.Errorf("Testing post-order iterator for empty tree. Wrong indication of hasNext. Expected %t. Actual %t.", false, iter.HasNext())
	}
}

func TestPostOrderIterator_HasNextOfNonEmptyTree(t *testing.T) {
	tree := GetSingleNodeTree()
	iter := NewPostOrderIterator(tree)

	if iter.HasNext() != true {
		t.Errorf("Testing post-order iterator for empty tree. Wrong indication of hasNext. Expected %t. Actual %t.", true, iter.HasNext())
	}
}

func TestPostOrderIterator_IterSingleTree(t *testing.T) {
	tree := GetSingleNodeTree()
	iter := NewPostOrderIterator(tree)
	examinePostOrderIterator(t, tree.root, iter)
}

func TestPostOrderIterator_IterAllLeftTree(t *testing.T) {
	tree := GetAllLeftNodeTree(5)
	iter := NewPostOrderIterator(tree)
	examinePostOrderIterator(t, tree.root, iter)
}

func TestPostOrderIterator_IterAllRightTree(t *testing.T) {
	tree := GetAllRightNodeTree(4)
	iter := NewPostOrderIterator(tree)
	examinePostOrderIterator(t, tree.root, iter)
}

func TestPostOrderIterator_IterFullTree(t *testing.T) {
	tree := GetFullTree(3)
	iter := NewPostOrderIterator(tree)
	examinePostOrderIterator(t, tree.root, iter)
}

func TestPostOrderIterator_IterRandomTree(t *testing.T) {
	tree := GetRandomTree(6)
	iter := NewPostOrderIterator(tree)
	examinePostOrderIterator(t, tree.root, iter)
}

func examinePostOrderIterator(t *testing.T, node *Node, iter Iterator) {
	examinePostOrderIteratorDfs(t, node, iter)
	if iter.HasNext() {
		t.Error("Iterator should have no more node to iterate.")
	}
}

func examinePostOrderIteratorDfs(t *testing.T, node *Node, iter Iterator) {
	if node == nil {
		return
	}

	examinePostOrderIteratorDfs(t, node.LeftChild, iter)
	examinePostOrderIteratorDfs(t, node.RightChild, iter)

	if !iter.HasNext() {
		t.Error("Iterator has no next node, but current node is not nil.")
		return
	}
	actualNode := iter.Next()
	if node.Value != actualNode.Value {
		t.Errorf("Iterator returns wrong node. Expected data value is %d. Actual data value is %d.", node.Value, actualNode.Value)
	}
}

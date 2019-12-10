package btree

import (
	"testing"

	"github.com/aiden0z/kit/base"
)

func assertValidTree(t *testing.T, tree *BTree, expectedSize int) {
	if actualValue, expectedValue := tree.size, expectedSize; actualValue != expectedValue {
		t.Errorf("Got %v expected %v for tree size", actualValue, expectedValue)
	}
}

func assertValidTreeNode(t *testing.T, node *Node, expectedEntries int, expectedChildren int, keys []base.Comparable, hasParent bool) {
	if actualValue, expectedValue := node.Parent != nil, hasParent; actualValue != expectedValue {
		t.Errorf("Got %v expected %v for hasParent", actualValue, expectedValue)
	}
	if actualValue, expectedValue := len(node.Entries), expectedEntries; actualValue != expectedValue {
		t.Errorf("Got %v expected %v for entries size", actualValue, expectedValue)
	}
	if actualValue, expectedValue := len(node.Children), expectedChildren; actualValue != expectedValue {
		t.Errorf("Got %v expected %v for children size", actualValue, expectedValue)
	}
	for i, key := range keys {
		if actualValue, expectedValue := node.Entries[i].Key, key; expectedValue.CompareTo(actualValue) != 0 {
			t.Errorf("Got %v expected %v for key", actualValue, expectedValue)
		}
	}
}

func TestBTreeGet1(t *testing.T) {
	tree := NewBTree(3)
	tree.Insert(base.Int(1), "a")
	tree.Insert(base.Int(2), "b")
	tree.Insert(base.Int(3), "c")
	tree.Insert(base.Int(4), "d")
	tree.Insert(base.Int(5), "e")
	tree.Insert(base.Int(6), "f")
	tree.Insert(base.Int(7), "g")

	tests := [][]interface{}{
		{base.Int(0), nil, false},
		{base.Int(1), "a", true},
		{base.Int(2), "b", true},
		{base.Int(3), "c", true},
		{base.Int(4), "d", true},
		{base.Int(5), "e", true},
		{base.Int(6), "f", true},
		{base.Int(7), "g", true},
		{base.Int(8), nil, false},
	}

	for _, test := range tests {
		comparable, _ := test[0].(base.Comparable)
		if value, found := tree.Get(comparable); value != test[1] || found != test[2] {
			t.Errorf("Got %v, %v expected %v, %v", value, found, test[1], test[2])
		}
	}

}

func TestGet2(t *testing.T) {

	tree := NewBTree(3)
	tree.Insert(base.Int(7), "g")
	tree.Insert(base.Int(9), "i")
	tree.Insert(base.Int(10), "j")
	tree.Insert(base.Int(6), "f")
	tree.Insert(base.Int(3), "c")
	tree.Insert(base.Int(4), "d")
	tree.Insert(base.Int(5), "e")
	tree.Insert(base.Int(8), "h")
	tree.Insert(base.Int(2), "b")
	tree.Insert(base.Int(1), "a")

	tests := [][]interface{}{
		{base.Int(0), nil, false},
		{base.Int(1), "a", true},
		{base.Int(2), "b", true},
		{base.Int(3), "c", true},
		{base.Int(4), "d", true},
		{base.Int(5), "e", true},
		{base.Int(6), "f", true},
		{base.Int(7), "g", true},
		{base.Int(8), "h", true},
		{base.Int(9), "i", true},
		{base.Int(10), "j", true},
		{base.Int(11), nil, false},
	}

	for _, test := range tests {
		comparable, _ := test[0].(base.Comparable)
		if value, found := tree.Get(comparable); value != test[1] || found != test[2] {
			t.Errorf("Got %v,%v expected %v,%v", value, found, test[1], test[2])
		}
	}
}

func TestBTreeInsert1(t *testing.T) {
	// https://upload.wikimedia.org/wikipedia/commons/3/33/B_tree_insertion_example.png
	tree := NewBTree(3)
	assertValidTree(t, tree, 0)

	tree.Insert(base.Int(1), 0)
	assertValidTree(t, tree, 1)
	assertValidTreeNode(t, tree.Root, 1, 0, []base.Comparable{base.Int(1)}, false)

	tree.Insert(base.Int(2), 1)
	assertValidTree(t, tree, 2)
	assertValidTreeNode(t, tree.Root, 2, 0, []base.Comparable{base.Int(1), base.Int(2)}, false)

	tree.Insert(base.Int(3), 2)
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(2)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(1)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []base.Comparable{base.Int(3)}, true)

	tree.Insert(base.Int(4), 2)
	assertValidTree(t, tree, 4)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(2)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(1)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 2, 0, []base.Comparable{base.Int(3), base.Int(4)}, true)

	tree.Insert(base.Int(5), 2)
	assertValidTree(t, tree, 5)
	assertValidTreeNode(t, tree.Root, 2, 3, []base.Comparable{base.Int(2), base.Int(4)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(1)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []base.Comparable{base.Int(3)}, true)
	assertValidTreeNode(t, tree.Root.Children[2], 1, 0, []base.Comparable{base.Int(5)}, true)

	tree.Insert(base.Int(6), 2)
	assertValidTree(t, tree, 6)
	assertValidTreeNode(t, tree.Root, 2, 3, []base.Comparable{base.Int(2), base.Int(4)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(1)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []base.Comparable{base.Int(3)}, true)
	assertValidTreeNode(t, tree.Root.Children[2], 2, 0, []base.Comparable{base.Int(5), base.Int(6)}, true)

	tree.Insert(base.Int(7), 2)
	assertValidTree(t, tree, 7)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(4)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 2, []base.Comparable{base.Int(2)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 2, []base.Comparable{base.Int(6)}, true)
	assertValidTreeNode(t, tree.Root.Children[0].Children[0], 1, 0, []base.Comparable{base.Int(1)}, true)
	assertValidTreeNode(t, tree.Root.Children[0].Children[1], 1, 0, []base.Comparable{base.Int(3)}, true)
	assertValidTreeNode(t, tree.Root.Children[1].Children[0], 1, 0, []base.Comparable{base.Int(5)}, true)
	assertValidTreeNode(t, tree.Root.Children[1].Children[1], 1, 0, []base.Comparable{base.Int(7)}, true)
}

func TestBTreePut2(t *testing.T) {
	tree := NewBTree(4)
	assertValidTree(t, tree, 0)

	tree.Insert(base.Int(0), 0)
	assertValidTree(t, tree, 1)
	assertValidTreeNode(t, tree.Root, 1, 0, []base.Comparable{base.Int(0)}, false)

	tree.Insert(base.Int(2), 2)
	assertValidTree(t, tree, 2)
	assertValidTreeNode(t, tree.Root, 2, 0, []base.Comparable{base.Int(0), base.Int(2)}, false)

	tree.Insert(base.Int(1), 1)
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 3, 0, []base.Comparable{base.Int(0), base.Int(1), base.Int(2)}, false)

	tree.Insert(base.Int(1), 1)
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 3, 0, base.NewIntComparableSlice([]int{0, 1, 2}), false)

	tree.Insert(base.Int(3), 3)
	assertValidTree(t, tree, 4)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(1)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(0)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 2, 0, []base.Comparable{base.Int(2), base.Int(3)}, true)

	tree.Insert(base.Int(4), 4)
	assertValidTree(t, tree, 5)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(1)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(0)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 3, 0, []base.Comparable{base.Int(2), base.Int(3), base.Int(4)}, true)

	tree.Insert(base.Int(5), 5)
	assertValidTree(t, tree, 6)
	assertValidTreeNode(t, tree.Root, 2, 3, []base.Comparable{base.Int(1), base.Int(3)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(0)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []base.Comparable{base.Int(2)}, true)
	assertValidTreeNode(t, tree.Root.Children[2], 2, 0, []base.Comparable{base.Int(4), base.Int(5)}, true)
}

func TestBTreeRemove1(t *testing.T) {
	// empty
	tree := NewBTree(3)
	tree.Remove(base.Int(1))
	assertValidTree(t, tree, 0)
}

func TestBTreeRemove2(t *testing.T) {
	// leaf node (no underflow)
	tree := NewBTree(3)
	tree.Insert(base.Int(1), nil)
	tree.Insert(base.Int(2), nil)

	tree.Remove(base.Int(1))
	assertValidTree(t, tree, 1)
	assertValidTreeNode(t, tree.Root, 1, 0, []base.Comparable{base.Int(2)}, false)

	tree.Remove(base.Int(2))
	assertValidTree(t, tree, 0)
}

func TestBTreeRemove4(t *testing.T) {
	// rotate left (underflow)
	tree := NewBTree(3)
	tree.Insert(base.Int(1), nil)
	tree.Insert(base.Int(2), nil)
	tree.Insert(base.Int(3), nil)
	tree.Insert(base.Int(4), nil)

	assertValidTree(t, tree, 4)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(2)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(1)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 2, 0, []base.Comparable{base.Int(3), base.Int(4)}, true)

	tree.Remove(base.Int(1))
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(3)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(2)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []base.Comparable{base.Int(4)}, true)
}

func TestBTreeRemove5(t *testing.T) {
	// rotate right (underflow)
	tree := NewBTree(3)
	tree.Insert(base.Int(1), nil)
	tree.Insert(base.Int(2), nil)
	tree.Insert(base.Int(3), nil)
	tree.Insert(base.Int(0), nil)

	assertValidTree(t, tree, 4)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(2)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 2, 0, []base.Comparable{base.Int(0), base.Int(1)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []base.Comparable{base.Int(3)}, true)

	tree.Remove(base.Int(3))
	assertValidTree(t, tree, 3)
	assertValidTreeNode(t, tree.Root, 1, 2, []base.Comparable{base.Int(1)}, false)
	assertValidTreeNode(t, tree.Root.Children[0], 1, 0, []base.Comparable{base.Int(0)}, true)
	assertValidTreeNode(t, tree.Root.Children[1], 1, 0, []base.Comparable{base.Int(2)}, true)
}

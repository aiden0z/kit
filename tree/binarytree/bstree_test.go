package binarytree

import (
	"fmt"
	"testing"

	"github.com/aiden0z/kit/base"
)

func TestBSTreeFind(t *testing.T) {

	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBSTreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	key := base.Int(5)

	node := btree.Find(key)

	if node == nil {
		t.Errorf("Btree Find work error, can not found equivalent node, %s", key)
	}

	if key.CompareTo(node.Element) != 0 {

		t.Errorf("Btree Find work error, not find the correct equivalent node")
	}

	// find no exist node
	key = base.Int(100)
	node = btree.Find(key)
	if node != nil {
		t.Errorf("Btree Find work error, find a non exist targe")
	}

}

func TestBSTreeFindNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	postOrder := base.NewIntComparableSlice([]int{1, 2, 5, 3, 8, 11, 10, 9, 6})

	bstree, err := NewBSTreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	key := base.Int(5)

	node := bstree.FindNonRecursive(key)

	if node == nil {
		t.Errorf("BSTree FindNonRecursive work error, can not found equivalent node, %s", key)
		return
	}

	if key.CompareTo(node.Element) != 0 {

		t.Errorf("Btree FindNonRecursive  work error, not find the correct equivalent node")
	}

	// find no exist node
	key = base.Int(100)
	node = bstree.FindNonRecursive(key)
	if node != nil {
		t.Errorf("Btree FindNonRecursive work error, find a non exist targe")
	}
}

func TestBSTreeFindMax(t *testing.T) {

	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBSTreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	max := base.Int(11)

	maxNode := btree.FindMax()

	if maxNode == nil {
		t.Error("Btree FindMax work error, can not find the max node.")
	}

	if max.CompareTo(maxNode.Element) != 0 {
		t.Error("Btree FindMax work error, not find the correct max node.")
	}

}

func TestBSTreeFindMaxNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBSTreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	max := base.Int(11)

	maxNode := btree.FindMaxNonRecursive()

	if maxNode == nil {
		t.Error("Btree FindMaxNonRecursive work error, can not find the max node.")
	}

	if max.CompareTo(maxNode.Element) != 0 {
		t.Error("Btree FindMaxNonRecursive  work error, not find the correct max node.")
	}

}

func TestBSTreeFindMin(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBSTreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	min := base.Int(1)

	minNode := btree.FindMin()

	if minNode == nil {
		t.Error("Btree FindMin work error, can not find the min node.")
	}

	if min.CompareTo(minNode.Element) != 0 {
		t.Error("Btree FindMin work error, not find the correct min node.")
	}

}

func TestBSTreeFindMinNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBSTreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	min := base.Int(1)

	minNode := btree.FindMinNonRecursive()

	if minNode == nil {
		t.Error("Btree FindMinNonRecursive work error, can not find the min node.")
	}

	if min.CompareTo(minNode.Element) != 0 {
		t.Error("Btree FindMinNonRecursive work error, not find the correct min node.")
	}

}

func TestBSTreeInsert(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBSTreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	key := base.Int(0)

	aBtree := btree.Insert(key)

	minNode := aBtree.FindMin()

	if key.CompareTo(minNode.Element) != 0 {
		t.Error("Btree Insert wrok error")
	}
}

func TestBSTreeInsertNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBSTreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)

	}

	fmt.Print(btree.HorizontalPretty().String())

	key := base.Int(7)

	aBtree := btree.InsertNonRecursive(key)

	fmt.Print(btree.HorizontalPretty().String())

	fmt.Print(btree.VerticalPretty().String())

	node := aBtree.Find(key)

	if key.CompareTo(node.Element) != 0 {
		t.Error("Btree InsertNonRecursive wrok error")
	}
}

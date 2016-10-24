package tree

import (
	"github.com/aiden0z/kit/base"
	"testing"
)

func TestNewBtreeWithPreInOrder(t *testing.T) {
	preOrder := base.NewIntComparableSlice([]int{7, 10, 4, 3, 1, 2, 8, 11})
	inOrder := base.NewIntComparableSlice([]int{4, 10, 3, 1, 7, 11, 8, 2})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result := btree.PreOrder()

	if len(result) == 0 {
		t.Error("build PRE-Order failed")
	}
	for i, v := range result {
		if preOrder[i] != v.Element {
			t.Error("build PRE-Order failed")
		}
	}

	result = btree.InOrder()
	if len(result) == 0 {
		t.Error("build IN-Order failed")
	}
	for i, v := range result {
		if inOrder[i] != v.Element {
			t.Error("build IN-Order failed")
		}
	}

	preOrder1 := base.NewRuneComparableSlice([]rune{'a', 'b', 'd', 'e', 'f', 'g', 'c'})
	inOrder1 := base.NewRuneComparableSlice([]rune{'d', 'e', 'b', 'g', 'f', 'a', 'c'})

	btree, err = NewBtreeWithInPreOrder(inOrder1, preOrder1)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result = btree.PreOrder()

	for i, v := range result {
		if preOrder1[i] != v.Element {
			t.Error("build PRE-Order failed")
		}
	}
}

func TestNewBtreeWithPostInOrder(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{4, 10, 3, 1, 7, 11, 8, 2})
	postOrder := base.NewIntComparableSlice([]int{4, 1, 3, 10, 11, 8, 2, 7})

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result := btree.PostOrder()

	for i, v := range result {
		if postOrder[i] != v.Element {
			t.Error("build POST-Order failed")
		}
	}

	result = btree.InOrder()
	for i, v := range result {
		if inOrder[i] != v.Element {
			t.Error("build IN-Order failed")
		}
	}

	postOrder1 := base.NewRuneComparableSlice([]rune{'e', 'd', 'g', 'f', 'b', 'c', 'a'})
	inOrder1 := base.NewRuneComparableSlice([]rune{'d', 'e', 'b', 'g', 'f', 'a', 'c'})

	btree, err = NewBtreeWithInPostOrder(inOrder1, postOrder1)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result = btree.PostOrder()

	for i, v := range result {
		if postOrder1[i] != v.Element {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtreePreOrderNonRecursive(t *testing.T) {
	preOrder := base.NewIntComparableSlice([]int{7, 10, 4, 3, 1, 2, 8, 11})
	inOrder := base.NewIntComparableSlice([]int{4, 10, 3, 1, 7, 11, 8, 2})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PreOrderNonRecursive()

	if len(result) == 0 {
		t.Error("build PRE-Order failed")
	}

	for i, v := range result {
		if preOrder[i] != v.Element {
			t.Error("build PRE-Order failed")
		}
	}
}

func TestBtreePreOrderMorris(t *testing.T) {
	preOrder := base.NewIntComparableSlice([]int{7, 10, 4, 3, 1, 2, 8, 11})
	inOrder := base.NewIntComparableSlice([]int{4, 10, 3, 1, 7, 11, 8, 2})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PreOrderMorris()

	if len(result) == 0 {
		t.Error("build PRE-Order failed")
	}

	for i, v := range result {
		if preOrder[i] != v.Element {
			t.Error("build PRE-Order failed")
		}
	}
}

func TestBtreeInOrderNonRecursive(t *testing.T) {
	preOrder := base.NewIntComparableSlice([]int{7, 10, 4, 3, 1, 2, 8, 11})
	inOrder := base.NewIntComparableSlice([]int{4, 10, 3, 1, 7, 11, 8, 2})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.InOrderNonRecursive()

	if len(result) == 0 {
		t.Error("build IN-Order failed")
	}

	for i, v := range result {
		if inOrder[i] != v.Element {
			t.Error("build IN-Order failed")
		}
	}
}

func TestBtreeInOrderMorris(t *testing.T) {
	preOrder := base.NewIntComparableSlice([]int{7, 10, 4, 3, 1, 2, 8, 11})
	inOrder := base.NewIntComparableSlice([]int{4, 10, 3, 1, 7, 11, 8, 2})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.InOrderMorris()

	if len(result) == 0 {
		t.Error("build IN-Order failed")
	}

	for i, v := range result {
		if inOrder[i] != v.Element {
			t.Error("build IN-Order failed")
		}
	}
}

func TestBtreePostOrderNonRecursive(t *testing.T) {
	postOrder := base.NewRuneComparableSlice([]rune{'e', 'd', 'g', 'f', 'b', 'c', 'a'})
	inOrder := base.NewRuneComparableSlice([]rune{'d', 'e', 'b', 'g', 'f', 'a', 'c'})

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursive()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i] != v.Element {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtreePostOrderMorris(t *testing.T) {

	postOrder := base.NewRuneComparableSlice([]rune{'e', 'd', 'g', 'f', 'b', 'c', 'a'})
	inOrder := base.NewRuneComparableSlice([]rune{'d', 'e', 'b', 'g', 'f', 'a', 'c'})

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderMorris()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i] != v.Element {
			t.Error("build POST-Order failed")
		}
	}

}

func TestBtreePostOrderNonRecursiveV1(t *testing.T) {
	postOrder := base.NewRuneComparableSlice([]rune{'e', 'd', 'g', 'f', 'b', 'c', 'a'})
	inOrder := base.NewRuneComparableSlice([]rune{'d', 'e', 'b', 'g', 'f', 'a', 'c'})

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursiveV1()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i] != v.Element {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtreePostOrderNonRecursiveV2(t *testing.T) {
	postOrder := base.NewRuneComparableSlice([]rune{'e', 'd', 'g', 'f', 'b', 'c', 'a'})
	inOrder := base.NewRuneComparableSlice([]rune{'d', 'e', 'b', 'g', 'f', 'a', 'c'})

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursiveV2()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i] != v.Element {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtree_LevelOrder(t *testing.T) {

	postOrder := base.NewRuneComparableSlice([]rune{'e', 'd', 'g', 'f', 'b', 'c', 'a'})
	inOrder := base.NewRuneComparableSlice([]rune{'d', 'e', 'b', 'g', 'f', 'a', 'c'})
	levelOrder := base.NewRuneComparableSlice([]rune{'a', 'b', 'c', 'd', 'f', 'e', 'g'})
	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.LevelOrder()

	if len(result) == 0 {
		t.Error("build Level-Order failed")
	}

	for i, v := range result {
		if levelOrder[i] != v.Element {
			t.Error("build Level-Order failed")
		}
	}

}

func TestBtree_Find(t *testing.T) {

	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err !=nil {
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

func TestBtree_FindNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	postOrder := base.NewIntComparableSlice([]int{1, 2, 5, 3, 8, 11, 10, 9, 6})

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err !=nil {
		t.Errorf("build btree failed %s", err)
	}
	key := base.Int(5)

	node := btree.FindNonRecursive(key)

	if node == nil {
		t.Errorf("Btree FindNonRecursive work error, can not found equivalent node, %s", key)
	}

	if key.CompareTo(node.Element) != 0 {

		t.Errorf("Btree FindNonRecursive  work error, not find the correct equivalent node")
	}

	// find no exist node
	key = base.Int(100)
	node = btree.FindNonRecursive(key)
	if node != nil {
		t.Errorf("Btree FindNonRecursive work error, find a non exist targe")
	}
}

func TestBtree_FindMax(t *testing.T) {

	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err !=nil {
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

func TestBtree_FindMaxNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err !=nil {
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

func TestBtree_FindMin(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err !=nil {
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

func TestBtree_FindMinNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err !=nil {
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

func TestBtree_Insert(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err !=nil {
		t.Errorf("build btree failed %s", err)
	}

	key := base.Int(0)

	aBtree := btree.Insert(key)

	minNode := aBtree.FindMin()

	if key.CompareTo(minNode.Element) != 0 {
		t.Error("Btree Insert wrok error")
	}
}

func TestBtree_InsertNonRecursive(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err !=nil {
		t.Errorf("build btree failed %s", err)
	}

	key := base.Int(7)

	aBtree := btree.InsertNonRecursive(key)

	node := aBtree.Find(key)

	if key.CompareTo(node.Element) != 0 {
		t.Error("Btree InsertNonRecursive wrok error")
	}
}

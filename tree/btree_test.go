package tree

import (
	"testing"
	"github.com/aiden0z/kit/base"
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

func TestBtreeLevelOrder(t *testing.T) {

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

func TestBtree_Depth(t *testing.T) {
	inOrder := base.NewIntComparableSlice([]int{1, 2, 3, 5, 6, 8, 9, 10, 11})
	preOrder := base.NewIntComparableSlice([]int{6, 3, 2, 1, 5, 9, 8, 10, 11})

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	if btree.Depth() != 4 {
		t.Error("Btree depth return incorrect value")

	}
}

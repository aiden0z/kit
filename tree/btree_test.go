package tree

import (
	"testing"
)

func TestNewBtreeWithPreInOrder(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result := btree.PreOrder()

	if len(result) == 0 {
		t.Error("build PRE-Order failed")
	}
	for i, v := range result {
		if preOrder[i].(int) != v.Element.(int) {
			t.Error("build PRE-Order failed")
		}
	}

	result = btree.InOrder()
	if len(result) == 0 {
		t.Error("build IN-Order failed")
	}
	for i, v := range result {
		if inOrder[i].(int) != v.Element.(int) {
			t.Error("build IN-Order failed")
		}
	}

	preOrder = []interface{}{'a', 'b', 'd', 'e', 'f', 'g', 'c'}
	inOrder = []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err = NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result = btree.PreOrder()

	for i, v := range result {
		if preOrder[i].(rune) != v.Element.(rune) {
			t.Error("build PRE-Order failed")
		}
	}
}

func TestNewBtreeWithPostInOrder(t *testing.T) {
	postOrder := []interface{}{4, 1, 3, 10, 11, 8, 2, 7}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result := btree.PostOrder()

	for i, v := range result {
		if postOrder[i].(int) != v.Element.(int) {
			t.Error("build POST-Order failed")
		}
	}

	result = btree.InOrder()
	for i, v := range result {
		if inOrder[i].(int) != v.Element.(int) {
			t.Error("build IN-Order failed")
		}
	}

	postOrder = []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder = []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err = NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result = btree.PostOrder()

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtreePreOrderNonRecursive(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PreOrderNonRecursive()

	if len(result) == 0 {
		t.Error("build PRE-Order failed")
	}

	for i, v := range result {
		if preOrder[i].(int) != v.Element.(int) {
			t.Error("build PRE-Order failed")
		}
	}
}

func TestBtreePreOrderMorris(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PreOrderMorris()

	if len(result) == 0 {
		t.Error("build PRE-Order failed")
	}

	for i, v := range result {
		if preOrder[i].(int) != v.Element.(int) {
			t.Error("build PRE-Order failed")
		}
	}
}

func TestBtreeInOrderNonRecursive(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.InOrderNonRecursive()

	if len(result) == 0 {
		t.Error("build IN-Order failed")
	}

	for i, v := range result {
		if inOrder[i].(int) != v.Element.(int) {
			t.Error("build IN-Order failed")
		}
	}
}

func TestBtreeInOrderMorris(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithInPreOrder(inOrder, preOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.InOrderMorris()

	if len(result) == 0 {
		t.Error("build IN-Order failed")
	}

	for i, v := range result {
		if inOrder[i].(int) != v.Element.(int) {
			t.Error("build IN-Order failed")
		}
	}
}

func TestBtreePostOrderNonRecursive(t *testing.T) {
	postOrder := []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder := []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursive()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtreePostOrderMorris(t *testing.T) {

	postOrder := []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder := []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderMorris()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build POST-Order failed")
		}
	}

}

func TestBtreePostOrderNonRecursiveV1(t *testing.T) {
	postOrder := []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder := []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursiveV1()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtreePostOrderNonRecursiveV2(t *testing.T) {
	postOrder := []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder := []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursiveV2()

	if len(result) == 0 {
		t.Error("build POST-Order failed")
	}

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build POST-Order failed")
		}
	}
}

func TestBtree_LevelOrder(t *testing.T) {

	postOrder := []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder := []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}
	levelOrder := []interface{}{'a', 'b', 'c', 'd', 'f', 'e', 'g'}
	btree, err := NewBtreeWithInPostOrder(inOrder, postOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.LevelOrder()

	if len(result) == 0 {
		t.Error("build Level-Order failed")
	}

	for i, v := range result {
		if levelOrder[i].(rune) != v.Element.(rune) {
			t.Error("build Level-Order failed")
		}
	}

}

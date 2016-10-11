package tree

import "testing"

func TestNewBtreeWithPreInOrder(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithPreInOrder(preOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result := btree.PreOrder()

	for i, v := range result {
		if preOrder[i].(int) != v.Element.(int) {
			t.Error("build pre order failed")
		}
	}

	result = btree.InOrder()
	for i, v := range result {
		if inOrder[i].(int) != v.Element.(int) {
			t.Error("build in order failed")
		}
	}

	preOrder = []interface{}{'a', 'b', 'd', 'e', 'f', 'g', 'c'}
	inOrder = []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err = NewBtreeWithPreInOrder(preOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result = btree.PreOrder()

	for i, v := range result {
		if preOrder[i].(rune) != v.Element.(rune) {
			t.Error("build pre order failed")
		}
	}
}

func TestNewBtreeWithPostInOrder(t *testing.T) {
	postOrder := []interface{}{4, 1, 3, 10, 11, 8, 2, 7}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithPostInOrder(postOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result := btree.PostOrder()

	for i, v := range result {
		if postOrder[i].(int) != v.Element.(int) {
			t.Error("build post order failed")
		}
	}

	result = btree.InOrder()
	for i, v := range result {
		if inOrder[i].(int) != v.Element.(int) {
			t.Error("build in order failed")
		}
	}

	postOrder = []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder = []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err = NewBtreeWithPostInOrder(postOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}
	result = btree.PostOrder()

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build post order failed")
		}
	}
}

func TestBtreePreOrderNonRecursive(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithPreInOrder(preOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PreOrderNonRecursive()

	for i, v := range result {
		if preOrder[i].(int) != v.Element.(int) {
			t.Error("build pre order failed")
		}
	}
}

func TestBtreeInOrderNonRecursive(t *testing.T) {
	preOrder := []interface{}{7, 10, 4, 3, 1, 2, 8, 11}
	inOrder := []interface{}{4, 10, 3, 1, 7, 11, 8, 2}

	btree, err := NewBtreeWithPreInOrder(preOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.InOrderNonRecursive()

	for i, v := range result {
		if inOrder[i].(int) != v.Element.(int) {
			t.Error("build pre order failed")
		}
	}
}

func TestBtreePostOrderNonRecursive(t *testing.T) {
	postOrder := []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder := []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err := NewBtreeWithPostInOrder(postOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursive()

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build post order failed")
		}
	}
}

func TestBtreePostOrderNonRecursiveV1(t *testing.T) {
	postOrder := []interface{}{'e', 'd', 'g', 'f', 'b', 'c', 'a'}
	inOrder := []interface{}{'d', 'e', 'b', 'g', 'f', 'a', 'c'}

	btree, err := NewBtreeWithPostInOrder(postOrder, inOrder)

	if err != nil {
		t.Errorf("build btree failed %s", err)
	}

	result := btree.PostOrderNonRecursiveV1()

	for i, v := range result {
		if postOrder[i].(rune) != v.Element.(rune) {
			t.Error("build post order failed")
		}
	}
}

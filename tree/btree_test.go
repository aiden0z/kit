package tree

import "testing"

func TestBtree(t *testing.T) {
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
}

package tree

import (
	"errors"
	"fmt"
)

type Btree TreeNode

// indexWithInOrder find the k's index in inOrder
func indexWithInOrder(k interface{}, inOrder []interface{}) int {
	for i, v := range inOrder {
		if v == k {
			return i
		}
	}
	return -1
}

// NewBtreeWithPreInOrder create a binary tree based on PRE and IN order
func NewBtreeWithPreInOrder(preOrder, inOrder []interface{}) (btree *Btree, err error) {

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				err = errors.New("invalid order sequence")
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()

	if err != nil {
		return
	}

	if len(preOrder) == 0 {
		return btree, nil
	}

	if len(preOrder) != len(inOrder) {
		err = fmt.Errorf("length of order sequence not equal")
		return
	}

	rootIndex := indexWithInOrder(preOrder[0], inOrder)

	btree = &Btree{
		Element: preOrder[0],
	}

	if node, e := NewBtreeWithPreInOrder(preOrder[1:rootIndex+1], inOrder[:rootIndex]); e != nil {
		return nil, e
	} else {
		btree.Left = (*TreeNode)(node)
	}

	if node, e := NewBtreeWithPreInOrder(preOrder[1+rootIndex:], inOrder[1+rootIndex:]); e != nil {
		return nil, e

	} else {
		btree.Right = (*TreeNode)(node)
	}
	return btree, nil
}

// NewBtreeWithPostInOrder create a binary tree based on POST and IN order
func NewBtreeWithPostInOrder(postOrder, inOrder []interface{}, offset int) *Btree {
	return nil

}

// PreOrder returh the pre order traversal
func (tree *Btree) PreOrder() (order []*TreeNode) {
	if tree != nil {
		order = append(order, (*TreeNode)(tree))
		order = append(order, (*Btree)(tree.Left).PreOrder()...)
		order = append(order, (*Btree)(tree.Right).PreOrder()...)
	}
	return order
}

// InOrder return the in order traversal
func (tree *Btree) InOrder() (order []*TreeNode) {
	if tree != nil {
		order = append(order, (*Btree)(tree.Left).InOrder()...)
		order = append(order, (*TreeNode)(tree))
		order = append(order, (*Btree)(tree.Right).InOrder()...)

	}
	return order
}

// PosrtOrder return the post order traversal
func (tree *Btree) PostOrder() (order []*TreeNode) {

	if tree != nil {
		order = append(order, (*Btree)(tree.Left).PostOrder()...)
		order = append(order, (*Btree)(tree.Right).PostOrder()...)
		order = append(order, (*TreeNode)(tree))

	}
	return order
}

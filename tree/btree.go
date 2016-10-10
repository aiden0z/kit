package tree

import (
	"errors"
	"fmt"
)

// Btree describe a binary tree
type Btree Node

// indexInSlice  find the k's index in slice
func indexInSlice(k interface{}, slice []interface{}) int {
	for i, v := range slice {
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
				fmt.Println(r)
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

	rootIndex := indexInSlice(preOrder[0], inOrder)

	btree = &Btree{
		Element: preOrder[0],
	}

	node, err := NewBtreeWithPreInOrder(preOrder[1:rootIndex+1], inOrder[:rootIndex])
	if err != nil {
		return nil, err
	}
	btree.Left = (*Node)(node)

	node, err = NewBtreeWithPreInOrder(preOrder[1+rootIndex:], inOrder[1+rootIndex:])
	if err != nil {
		return nil, err
	}

	btree.Right = (*Node)(node)
	return btree, nil
}

// NewBtreeWithPostInOrder create a binary tree based on POST and IN order
func NewBtreeWithPostInOrder(postOrder, inOrder []interface{}) (btree *Btree, err error) {

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

	if len(postOrder) == 0 {
		return btree, nil
	}

	if len(postOrder) != len(inOrder) {
		err = fmt.Errorf("length of order sequence not equal")
		return
	}

	rootIndex := indexInSlice(postOrder[len(postOrder)-1], inOrder)

	btree = &Btree{
		Element: postOrder[len(postOrder)-1],
	}

	node, err := NewBtreeWithPostInOrder(postOrder[:rootIndex], inOrder[:rootIndex])
	if err != nil {
		return nil, err
	}
	btree.Left = (*Node)(node)

	node, err = NewBtreeWithPostInOrder(postOrder[rootIndex:len(postOrder)-1], inOrder[1+rootIndex:])
	if err != nil {
		return nil, err
	}

	btree.Right = (*Node)(node)
	return btree, nil

}

// PreOrder returh the pre order traversal
func (tree *Btree) PreOrder() (order []*Node) {
	if tree != nil {
		order = append(order, (*Node)(tree))
		order = append(order, (*Btree)(tree.Left).PreOrder()...)
		order = append(order, (*Btree)(tree.Right).PreOrder()...)
	}
	return order
}

// InOrder return the in order traversal
func (tree *Btree) InOrder() (order []*Node) {
	if tree != nil {
		order = append(order, (*Btree)(tree.Left).InOrder()...)
		order = append(order, (*Node)(tree))
		order = append(order, (*Btree)(tree.Right).InOrder()...)

	}
	return order
}

// PostOrder return the post order traversal
func (tree *Btree) PostOrder() (order []*Node) {

	if tree != nil {
		order = append(order, (*Btree)(tree.Left).PostOrder()...)
		order = append(order, (*Btree)(tree.Right).PostOrder()...)
		order = append(order, (*Node)(tree))

	}
	return order
}

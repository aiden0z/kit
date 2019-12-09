//        Binary Search Tree
//
//                 6
//              /    \
//            /        \
//           3          9
//         /   \      /   \
//        2     5    8     10
//      /                     \
//     1                       11
//

// PRE-Order:   6, 3, 2, 1, 5, 9, 8, 10, 11
// IN-Order:    1, 2, 3, 5, 6, 8, 9, 10, 11
// POST-Order:  1, 2, 5, 3, 8, 11, 10, 9, 6

package binarytree

import (
	"bytes"

	"github.com/aiden0z/kit/base"
)

// BSTree present a binary search tree
type BSTree Btree

// NewBSTreeWithInPreOrder create a binary search tree based on PRE and IN order
func NewBSTreeWithInPreOrder(inOrder, preOrder []base.Comparable) (bstree *BSTree, err error) {
	var btree *Btree
	btree, err = NewBtreeWithInPreOrder(inOrder, preOrder)
	if err != nil {
		return nil, err
	}
	return (*BSTree)(btree), nil
}

// NewBSTreeWithInPostOrder create a binary search tree based on POST and IN order
func NewBSTreeWithInPostOrder(inOrder, postOrder []base.Comparable) (bstree *BSTree, err error) {
	var btree *Btree
	btree, err = NewBtreeWithInPostOrder(inOrder, postOrder)
	if err != nil {
		return nil, err
	}
	return (*BSTree)(btree), nil
}

// Find the specified node
func (tree *BSTree) Find(o base.Comparable) (node *BSTree) {
	if tree == nil {
		return
	}

	current := tree

	result := current.Element.CompareTo(o)

	if result < 0 {

		return (*BSTree)(current.Right).Find(o)
	} else if result > 0 {
		return (*BSTree)(current.Left).Find(o)
	} else {
		node = current
	}

	return
}

// FindNonRecursive find the code without recursive.
func (tree *BSTree) FindNonRecursive(o base.Comparable) (node *BSTree) {
	if tree == nil {
		return
	}

	node = tree

	var result int

	for node != nil {
		result = node.Element.CompareTo(o)

		if result == 0 {
			return
		} else if result > 0 {
			node = (*BSTree)(node.Left)
		} else {
			node = (*BSTree)(node.Right)
		}
	}

	return nil
}

// FindMin return minimum node
func (tree *BSTree) FindMin() (node *BSTree) {
	if tree == nil {
		return nil
	}

	node = tree

	if node.Left == nil {
		return node
	}

	return (*BSTree)(node.Left).FindMin()
}

// FindMinNonRecursive return minimum node
func (tree *BSTree) FindMinNonRecursive() (node *BSTree) {
	if tree == nil {
		return nil
	}

	node = tree
	for node.Left != nil {
		node = (*BSTree)(node.Left)
	}
	return
}

// FindMax return maximum node
func (tree *BSTree) FindMax() (node *BSTree) {

	if tree == nil {
		return nil
	}

	node = tree

	if node.Right == nil {
		return
	}

	return (*BSTree)(node.Right).FindMax()
}

// FindMaxNonRecursive return maximum node
func (tree *BSTree) FindMaxNonRecursive() (node *BSTree) {
	if tree == nil {
		return nil
	}

	node = tree

	for node.Right != nil {
		node = (*BSTree)(node.Right)
	}
	return
}

// Insert a value and return a inserted tree
func (tree *BSTree) Insert(o base.Comparable) (node *BSTree) {

	node = tree

	if node == nil {
		node = (*BSTree)(&Btree{Element: o})
	} else if (o.CompareTo(tree.Element)) < 0 {
		node.Left = (*Btree)((*BSTree)(node.Left).Insert(o))

	} else if (o.CompareTo(tree.Element)) > 0 {
		node.Right = (*Btree)((*BSTree)(node.Right).Insert(o))
	}

	return
}

// InsertNonRecursive insert a value non-recursive
func (tree *BSTree) InsertNonRecursive(o base.Comparable) (node *BSTree) {

	node = tree

	if node == nil {
		node = (*BSTree)(&Btree{Element: o})
		return
	}

	current := (*Btree)(node)

	for {
		if o.CompareTo(current.Element) < 0 {
			if current.Left == nil {
				current.Left = &Btree{Element: o}
				return
			}
			current = current.Left

		} else if o.CompareTo(current.Element) > 0 {
			if current.Right == nil {
				current.Right = &Btree{Element: o}
				return
			}
			current = current.Right

		} else {
			return
		}
	}
}

// VerticalPretty print the tree in vertical format.
func (tree *BSTree) VerticalPretty() *bytes.Buffer {
	return (*Btree)(tree).VerticalPretty()
}

// HorizontalPretty print the tree in horizontal format.
func (tree *BSTree) HorizontalPretty() *bytes.Buffer {
	return (*Btree)(tree).HorizontalPretty()
}

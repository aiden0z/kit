//            Binary tree
//
//                      a
//                    /   \
//                  b       c
//               /     \
//             d        f
//               \    /
//                e  g
//
// PRE-Order:  a, b, d, e, f, g, c
// IN-Order:   d, e, b, g, f, a, c
// POST-Order: e, d, g, f, b, c, a

package tree

import (
	"errors"
	"github.com/aiden0z/kit/stack"
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

// NewBtreeWithInPreOrder create a binary tree based on PRE and IN order
func NewBtreeWithInPreOrder(inOrder, preOrder []interface{}) (btree *Btree,
err error) {

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
		err = errors.New("length of order sequence not equal")
		return
	}

	rootIndex := indexInSlice(preOrder[0], inOrder)

	btree = &Btree{
		Element: preOrder[0],
	}

	node, err := NewBtreeWithInPreOrder(inOrder[:rootIndex], preOrder[1:rootIndex + 1])
	if err != nil {
		return nil, err
	}
	btree.Left = (*Node)(node)

	node, err = NewBtreeWithInPreOrder(inOrder[1 + rootIndex:], preOrder[1 + rootIndex:])
	if err != nil {
		return nil, err
	}

	btree.Right = (*Node)(node)
	return btree, nil
}

// NewBtreeWithInPostOrder create a binary tree based on POST and IN order
func NewBtreeWithInPostOrder(inOrder, postOrder []interface{}) (btree *Btree,
err error) {

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
		err = errors.New("length of order sequence not equal")
		return
	}

	rootIndex := indexInSlice(postOrder[len(postOrder) - 1], inOrder)

	btree = &Btree{
		Element: postOrder[len(postOrder) - 1],
	}

	node, err := NewBtreeWithInPostOrder(inOrder[:rootIndex], postOrder[:rootIndex])
	if err != nil {
		return nil, err
	}
	btree.Left = (*Node)(node)

	node, err = NewBtreeWithInPostOrder(inOrder[1 + rootIndex:], postOrder[rootIndex:len (postOrder) - 1])
	if err != nil {
		return nil, err
	}

	btree.Right = (*Node)(node)
	return btree, nil

}

// PreOrder return the PRE order traversal
func (tree *Btree) PreOrder() (order []*Node) {
	if tree != nil {
		order = append(order, (*Node)(tree))
		order = append(order, (*Btree)(tree.Left).PreOrder()...)
		order = append(order, (*Btree)(tree.Right).PreOrder()...)
	}
	return order
}

// PreOrderNonRecursive return the PRE order traversal
func (tree *Btree) PreOrderNonRecursive() (order []*Node) {
	if tree == nil {
		return
	}

	node := (*Node)(tree)
	s := stack.NewStack(10)

	for node != nil || !s.IsEmpty() {
		for node != nil {
			order = append(order, node)
			s.Push(node)
			node = node.Left
		}

		if !s.IsEmpty() {
			node = s.Pop().(*Node)
			node = node.Right
		}
	}

	return order
}

// PreOrderMorris return the IN order traversal based on threaded binary tree
// In In-Order, the right node of current node's predecessor always is nil
func (tree *Btree) PreOrderMorris() (order []*Node) {

	current := (*Node)(tree)
	var preNode *Node

	for current != nil {
		if current.Left == nil {
			order = append(order, current)
			current = current.Right
		} else {
			// find the current node's predecessor in In-Order
			preNode = current.Left
			for preNode.Right != nil && preNode.Right != current {
				preNode = preNode.Right
			}

			if preNode.Right == nil {
				// first time visit, link the right node of current node's predecessor to current node
				order = append(order, current)
				preNode.Right = current
				current = current.Left
			} else {
				// second time visit, remove the link set in first time visit
				preNode.Right = nil
				current = current.Right
			}
		}

	}

	return
}

// InOrder return the IN order traversal
func (tree *Btree) InOrder() (order []*Node) {
	if tree != nil {
		order = append(order, (*Btree)(tree.Left).InOrder()...)
		order = append(order, (*Node)(tree))
		order = append(order, (*Btree)(tree.Right).InOrder()...)

	}
	return order
}

// InOrderNonRecursive return the IN order traversal
func (tree *Btree) InOrderNonRecursive() (order []*Node) {
	if tree == nil {
		return
	}

	node := (*Node)(tree)
	s := stack.NewStack(10)

	for node != nil || !s.IsEmpty() {
		for node != nil {
			s.Push(node)
			node = node.Left
		}

		if !s.IsEmpty() {
			node = s.Pop().(*Node)
			order = append(order, node)
			node = node.Right
		}
	}

	return order
}

// In In-Order, the right node of current node's predecessor always is nil
// In In-order traversal, start from root node
// 1. if current node has left child then find its in-order predecessor and
// make root as right child of it and move left of root. [ to find
// predecessor, find the max right element in its left subtree ]
// 2. if current node don't have left child , then print data and move right.
//
// The main thing which should be observe is that while performing step 1,
// we'll reach a point where predecessor right child is itself current node,
// this only happen when whole left child turned off and we start printing data from there.
//
// InOrderMorris return the IN order traversal based on threaded binary tree
func (tree *Btree) InOrderMorris() (order []*Node) {

	current := (*Node)(tree)
	var preNode *Node

	for current != nil {
		if current.Left == nil {
			order = append(order, current)
			current = current.Right
		} else {
			// find the current node's predecessor in In-Order
			preNode = current.Left
			for preNode.Right != nil && preNode.Right != current {
				preNode = preNode.Right
			}

			if preNode.Right == nil {
				// first time visit, link the right node of current node's predecessor to current node
				preNode.Right = current
				current = current.Left
			} else {
				// second time visit, remove the link set in first time visit, then output the node
				preNode.Right = nil
				order = append(order, current)
				current = current.Right
			}
		}

	}

	return
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

// PostOrderMorris return the POST order traversal based on threaded binary tree
func (tree *Btree) PostOrderMorris() (order []*Node) {

	dummyRoot := &Node{}
	// why link tree to the dummyRoot.Left ?
	// Because if we assume there is no right child of root then print
	// left child and then root become POST-order traversal
	dummyRoot.Left = (*Node)(tree)

	var preNode, first, middle, last *Node
	current := dummyRoot

	for (current != nil ) {
		if current.Left == nil {
			current = current.Right
		} else {
			// current has a left child, it also has a predecessor
			// find the current's predecessor in IN-order
			preNode = current.Left
			for (preNode.Right != nil && preNode.Right != current) {
				preNode = preNode.Right
			}

			// link the right child of predecessor to current
			// when predecessor found for first time
			if preNode.Right == nil {
				preNode.Right = current
				current = current.Left

			} else {
				// predecessor found second time
				// reverse the right references in chain from preNode to current node
				first = current
				middle = current.Left
				for (middle != current) {
					last = middle.Right
					middle.Right = first
					first = middle
					middle = last

				}

				// visit the nodes from preNode to current node
				// again reverse this right references from
				// preNode to current node
				first = current
				middle = preNode
				for (middle != current) {
					order = append(order, middle)
					last = middle.Right
					middle.Right = first
					first = middle
					middle = last
				}

				preNode.Right = nil
				current = current.Right

			}

		}

	}

	return
}

// PostOrderNonRecursive return the post order traversal
func (tree *Btree) PostOrderNonRecursive() (order []*Node) {
	if tree == nil {
		return
	}

	current := (*Node)(tree)

	var lastVisited *Node

	s := stack.NewStack(10)

	// push all left child into the stack firstly
	for current != nil {
		s.Push(current)
		current = current.Left
	}

	for !s.IsEmpty() {
		current = s.Peek().(*Node)
		// hand the root node
		if current.Right == nil || current.Right == lastVisited {
			order = append(order, current)
			lastVisited = current
			s.Pop()
		} else {
			// handle the right tree
			current = current.Right
			for current != nil {
				s.Push(current)
				current = current.Left
			}
		}
	}
	return
}

// PostOrderNonRecursiveV1 return the post order traversal
func (tree *Btree) PostOrderNonRecursiveV1() (order []*Node) {
	if tree == nil {
		return
	}

	current := (*Node)(tree)

	var lastVisited *Node

	s := stack.NewStack(10)

	// push the root node into stack firstly
	s.Push(current)

	for !s.IsEmpty() {
		current = s.Peek().(*Node)

		// root node
		if (current.Left == nil && current.Right == nil) || (lastVisited != nil && (lastVisited == current.Left || lastVisited == current.Right)) {
			order = append(order, current)
			s.Pop()
			lastVisited = current
		} else {
			// then push the right tree
			if current.Right != nil {
				s.Push(current.Right)
			}
			// pus the left tree
			if current.Left != nil {
				s.Push(current.Left)

			}
		}
	}
	return
}

// PostOrderNonRecursiveV2 return the post order traversal
func (tree *Btree) PostOrderNonRecursiveV2() (order []*Node) {

	if tree == nil {
		return
	}

	s1 := stack.NewStack(10)
	s2 := stack.NewStack(10)

	current := (*Node)(tree)
	s1.Push(current)

	// push order: root node -> left tree -> right tree
	for !s1.IsEmpty() {

		current = s1.Pop().(*Node)
		s2.Push(current)

		if current.Left != nil {
			s1.Push(current.Left)

		}

		if current.Right != nil {
			s1.Push(current.Right)
		}
	}

	// reverse the s1, right tree -> left tree -> root node, get the POST traversal order
	for !s2.IsEmpty() {

		order = append(order, s2.Pop().(*Node))
	}

	return
}

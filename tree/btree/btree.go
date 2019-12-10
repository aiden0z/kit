// Package btree implements a B tree.
// According to Knuth's definition, a B-tree of order m is a tree which
// satisfies the following properties:
// - Every node has at most m children.
// - Every non-leaf node (except root) has at least ⌈m/2⌉ children.
// - The root has at least two children if it is not a leaf node.
// - A non-leaf node with k children contains k−1 keys.
// - All leaves appear in the same level
// Reference https://github.com/emirpasic/gods/blob/master/trees/btree/btree.go
package btree

import (
	"fmt"

	"github.com/aiden0z/kit/base"
)

// BTree describe a b-tree.
type BTree struct {
	Root *Node
	size int // Total number of keys in the tree
	m    int // Maximum number of children of a node
}

// Node describe the tree node.
type Node struct {
	Parent   *Node
	Entries  []*Entry // The keys in node
	Children []*Node  // Children nodes
}

// Entry describe the keys in b-tree node.
type Entry struct {
	Key   base.Comparable
	Value interface{}
}

// NewBTree return a B-tree, order must greate than 2.
func NewBTree(order int) *BTree {
	return &BTree{
		m: order,
	}
}

func (entry *Entry) String() string {
	return fmt.Sprintf("%v", entry.Key)
}

// search key in node.
func (node *Node) search(key base.Comparable) (index int, found bool) {
	low, high := 0, len(node.Entries)-1
	var mid int

	for low <= high {
		mid = (high + low) / 2
		compare := key.CompareTo(node.Entries[mid].Key)
		if compare > 0 {
			low = mid + 1
		} else if compare < 0 {
			high = mid - 1
		} else {
			return mid, true
		}
	}
	return low, false
}

func (node *Node) isLeaf() bool {
	return len(node.Children) == 0
}

func (node *Node) height() int {
	hight := 0

	for ; node != nil; node = node.Children[0] {
		hight++
		if len(node.Children) == 0 {
			break
		}
	}
	return hight
}

func (node *Node) leftSibling(key base.Comparable) (*Node, int) {
	if node.Parent != nil {
		index, _ := node.Parent.search(key)
		index--
		if index >= 0 && index < len(node.Parent.Children) {
			return node.Parent.Children[index], index
		}
	}
	return nil, -1
}

func (node *Node) rightSibling(key base.Comparable) (*Node, int) {
	if node.Parent != nil {
		index, _ := node.Parent.search(key)
		index++
		if index < len(node.Parent.Children) {
			return node.Parent.Children[index], index
		}
	}

	return nil, -1
}

func (node *Node) left() *Node {
	current := node

	for {
		if current.isLeaf() {
			return current
		}
		current = current.Children[0]
	}
}

func (node *Node) right() *Node {
	current := node
	for {
		if current.isLeaf() {
			return current
		}

		current = current.Children[len(current.Children)-1]
	}
}

func (node *Node) prependChildrenFromNode(fromNode *Node) {
	children := append([]*Node{}, fromNode.Children...)
	node.Children = append(children, node.Children...)
	setParent(fromNode.Children, node)
}

func (node *Node) appendChilrenFromNode(fromNode *Node) {
	node.Children = append(node.Children, fromNode.Children...)
	setParent(fromNode.Children, node)
}

func (node *Node) deleteEntry(index int) {
	copy(node.Entries[index:], node.Entries[index+1:])
	node.Entries[len(node.Entries)-1] = nil
	node.Entries = node.Entries[:len(node.Entries)-1]
}

func (node *Node) deleteChild(index int) {
	if index >= len(node.Children) {
		return
	}

	copy(node.Children[index:], node.Children[index+1:])
	node.Children[len(node.Children)-1] = nil
	node.Children = node.Children[:len(node.Children)-1]
}

func (tree *BTree) maxChildren() int {
	return tree.m
}

func (tree *BTree) minChildren() int {
	return (tree.m + 1) / 2
}

func (tree *BTree) maxEntries() int {
	return tree.maxChildren() - 1
}

func (tree *BTree) minEntries() int {
	return tree.minChildren() - 1
}

func (tree *BTree) middle() int {
	return (tree.m - 1) / 2
}

func (tree *BTree) isFull(node *Node) bool {
	return len(node.Entries) == tree.maxEntries()
}

func (tree *BTree) shouldSplit(node *Node) bool {
	return len(node.Entries) > tree.maxEntries()
}

func (tree *BTree) searchRecursive(startNode *Node, key base.Comparable) (node *Node, index int, found bool) {
	if tree.Empty() {
		return nil, -1, false
	}

	node = startNode
	for {
		index, found = node.search(key)
		if found {
			return node, index, true
		}

		if node.isLeaf() {
			return nil, -1, false
		}

		node = node.Children[index]
	}
}

func (tree *BTree) split(node *Node) {

	if !tree.shouldSplit(node) {
		return
	}

	if node == tree.Root {
		tree.splitRoot()
		return
	}

	tree.splitNonRoot(node)
}

func (tree *BTree) splitNonRoot(node *Node) {
	middle := tree.middle()
	parent := node.Parent

	left := &Node{Entries: append([]*Entry{}, node.Entries[:middle]...), Parent: parent}
	right := &Node{Entries: append([]*Entry{}, node.Entries[middle+1:]...), Parent: parent}

	// move children from the node to be split into left and right node
	if !node.isLeaf() {
		left.Children = append([]*Node{}, node.Children[:middle+1]...)
		right.Children = append([]*Node{}, node.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	insertPosition, _ := parent.search(node.Entries[middle].Key)

	// insert middle key to parent
	parent.Entries = append(parent.Entries, nil)
	copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
	parent.Entries[insertPosition] = node.Entries[middle]

	// set child left of inserted key in parent to the created left node
	parent.Children[insertPosition] = left

	// set child right of inserted key in parent to the created right node
	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	tree.split(parent)
}

func (tree *BTree) splitRoot() {
	middle := tree.middle()

	left := &Node{Entries: append([]*Entry{}, tree.Root.Entries[:middle]...)}
	right := &Node{Entries: append([]*Entry{}, tree.Root.Entries[middle+1:]...)}

	// move children from node to be split into left and right nodes
	if !tree.Root.isLeaf() {
		left.Children = append([]*Node{}, tree.Root.Children[:middle+1]...)
		right.Children = append([]*Node{}, tree.Root.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	// root is a node with none entry and two children (left and right)
	newRoot := &Node{
		Entries:  []*Entry{tree.Root.Entries[middle]},
		Children: []*Node{left, right},
	}
	left.Parent = newRoot
	right.Parent = newRoot
	tree.Root = newRoot
}

func (tree *BTree) insert(node *Node, entry *Entry) (inserted bool) {
	if node.isLeaf() {
		return tree.insertToLeaf(node, entry)
	}
	return tree.insertToInternal(node, entry)
}

func (tree *BTree) insertToLeaf(node *Node, entry *Entry) (inserted bool) {
	insertPosition, found := node.search(entry.Key)

	// update
	if found {
		node.Entries[insertPosition] = entry
		return false
	}

	// insert entry's key in the middle of the node
	node.Entries = append(node.Entries, nil)
	copy(node.Entries[insertPosition+1:], node.Entries[insertPosition:])
	node.Entries[insertPosition] = entry
	tree.split(node)
	return true
}

func (tree *BTree) insertToInternal(node *Node, entry *Entry) (inserted bool) {
	insertPosition, found := node.search(entry.Key)

	// update
	if found {
		node.Entries[insertPosition] = entry
		return false
	}
	return tree.insert(node.Children[insertPosition], entry)
}

// delete deletes an entry in node at entries' index
func (tree *BTree) delete(node *Node, index int) {
	// deleteing from a leaf node
	if node.isLeaf() {
		deletedKey := node.Entries[index].Key
		node.deleteEntry(index)
		tree.rebalance(node, deletedKey)
		if len(tree.Root.Entries) == 0 {
			tree.Root = nil
		}

		return
	}

	// deleting from an internal node
	leftLargestNode := node.Children[index].right()
	leftLargestEntryIndex := len(leftLargestNode.Entries) - 1
	node.Entries[index] = leftLargestNode.Entries[leftLargestEntryIndex]
	deletedKey := leftLargestNode.Entries[leftLargestEntryIndex].Key
	leftLargestNode.deleteEntry(leftLargestEntryIndex)
	tree.rebalance(leftLargestNode, deletedKey)
}

// rebalance rebalances the tree after deletion.
// Note that we first delete the entry and then call rebalance, thus the passed
// deleted key as reference.
func (tree *BTree) rebalance(node *Node, deletedKey base.Comparable) {
	// check if rebalancing is required
	if node == nil || len(node.Entries) >= tree.minEntries() {
		return
	}

	// try to borrow from left sibling
	leftSibling, leftSiblingIndex := node.leftSibling(deletedKey)
	if leftSibling != nil && len(leftSibling.Entries) > tree.minEntries() {
		// rorate right
		node.Entries = append([]*Entry{node.Parent.Entries[leftSiblingIndex]}, node.Entries...)
		node.Parent.Entries[leftSiblingIndex] = leftSibling.Entries[len(leftSibling.Entries)-1]
		leftSibling.deleteEntry(len(leftSibling.Entries) - 1)
		if !leftSibling.isLeaf() {
			leftSiblingRightMostChild := leftSibling.Children[len(leftSibling.Children)-1]
			leftSiblingRightMostChild.Parent = node
			node.Children = append([]*Node{leftSiblingRightMostChild}, node.Children...)
			leftSibling.deleteChild(len(leftSibling.Children) - 1)
		}
		return
	}

	// try to borrow from right sibling
	rightSibling, rightSiblingIndex := node.rightSibling(deletedKey)
	if rightSibling != nil && len(rightSibling.Entries) > tree.minEntries() {
		// rotate left
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1])
		node.Parent.Entries[rightSiblingIndex-1] = rightSibling.Entries[0]
		rightSibling.deleteEntry(0)
		if !rightSibling.isLeaf() {
			rightSiblingLeftMostChild := rightSibling.Children[0]
			rightSiblingLeftMostChild.Parent = node
			node.Children = append(node.Children, rightSiblingLeftMostChild)
			rightSibling.deleteChild(0)
		}
		return
	}

	// merge with siblings
	if rightSibling != nil {
		// merge with right sibling
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1])
		node.Entries = append(node.Entries, rightSibling.Entries...)
		deletedKey = node.Parent.Entries[rightSiblingIndex-1].Key
		node.Parent.deleteEntry(rightSiblingIndex - 1)
		node.appendChilrenFromNode(node.Parent.Children[rightSiblingIndex])
		node.Parent.deleteChild(rightSiblingIndex)
	} else if leftSibling != nil {
		// merge with left sibling
		entries := append([]*Entry{}, leftSibling.Entries...)
		entries = append(entries, node.Parent.Entries[leftSiblingIndex])
		node.Entries = append(entries, node.Entries...)
		deletedKey = node.Parent.Entries[leftSiblingIndex].Key
		node.Parent.deleteEntry(leftSiblingIndex)
		node.prependChildrenFromNode(node.Parent.Children[leftSiblingIndex])
		node.Parent.deleteChild(leftSiblingIndex)
	}

	// make the merged node the root if its parent was root and the root is empty
	if node.Parent == tree.Root && len(tree.Root.Entries) == 0 {
		tree.Root = node
		node.Parent = nil
		return
	}

	// parent might underflow, so try to rebalance if necessary
	tree.rebalance(node.Parent, deletedKey)
}

func setParent(nodes []*Node, parent *Node) {
	for _, node := range nodes {
		node.Parent = parent
	}
}

// Clear removes all nodes from tree.
func (tree *BTree) Clear() {
	tree.Root = nil
	tree.size = 0
}

// Empty return true if three does not contains any nodes.
func (tree *BTree) Empty() bool {
	return tree.size == 0
}

// Size returns the number of nodes in the tree.
func (tree *BTree) Size() int {
	return tree.size
}

// Height returns height of the tree.
func (tree *BTree) Height() int {
	return tree.Root.height()
}

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *BTree) Left() *Node {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.left()
}

// Right return the right-most (max) node or nil if tree is empty.
func (tree *BTree) Right() *Node {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.right()
}

// Insert the key, value entry.
func (tree *BTree) Insert(key base.Comparable, value interface{}) {
	entry := &Entry{Key: key, Value: value}

	if tree.Root == nil {
		tree.Root = &Node{Entries: []*Entry{entry}, Children: []*Node{}}
		tree.size++
		return
	}

	if tree.insert(tree.Root, entry) {
		tree.size++
	}
}

// Get searches the node in tree by key and returns th its value or nil if key is not
// found in tree.
func (tree *BTree) Get(key base.Comparable) (value interface{}, found bool) {
	node, index, found := tree.searchRecursive(tree.Root, key)
	if found {
		return node.Entries[index].Value, true
	}
	return nil, false
}

// Remove remove the node from the tree by key.
func (tree *BTree) Remove(key base.Comparable) {
	node, index, found := tree.searchRecursive(tree.Root, key)
	if found {
		tree.delete(node, index)
		tree.size--
	}
}

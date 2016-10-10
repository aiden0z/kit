package tree

// Node describe the tree node
type Node struct {
	Element interface{}
	Left    *Node
	Right   *Node
}

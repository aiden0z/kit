package tree

import (
	"github.com/aiden0z/kit/base"
)



// Node describe the tree node
type Node struct {
	Element base.Comparable
	Left    *Node
	Right   *Node
}

package binarytree

import (
	"github.com/aiden0z/kit/base"
)

// Btree describe the tree node
type Btree struct {
	Element base.Comparable
	Left    *Btree
	Right   *Btree
}

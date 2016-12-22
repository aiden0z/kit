package tree

import (
	"github.com/aiden0z/kit/base"
)



// Node describe the tree node
type Btree struct {
	Element base.Comparable
	Left    *Btree
	Right   *Btree
}


package binarytree

import (
	"bytes"
	"fmt"
	"math"

	"github.com/aiden0z/kit/base"
)

// indexInSlice  find the k's index in slice
func indexInSlice(k base.Comparable, slice []base.Comparable) int {
	for i, v := range slice {
		if v == k {
			return i
		}
	}
	return -1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isAllNodesNull(nodes []*Btree) bool {
	for _, node := range nodes {
		if node != nil {
			return false
		}
	}
	return true

}

func whitespaceN(count int) string {
	if count == 0 {
		return ""
	}

	return fmt.Sprintf(fmt.Sprintf("%%%ds", count), " ")
}

func vertical(nodes []*Btree, level, depth int) (buffer *bytes.Buffer) {

	buffer = new(bytes.Buffer)

	if len(nodes) == 0 || isAllNodesNull(nodes) {
		return
	}

	floor := depth - level
	// 边的数量
	edges := int(math.Pow(2, float64((maxInt(floor-1, 0)))))
	// 行首空格数
	heads := int(math.Pow(2, float64(floor))) - 1
	// 间隔空格数
	space := int(math.Pow(2, float64(floor+1))) - 1

	buffer.WriteString(whitespaceN(heads))

	newNodes := make([]*Btree, 0)
	for _, node := range nodes {
		if node != nil {
			buffer.WriteString(node.Element.String())
			newNodes = append(newNodes, node.Left)
			newNodes = append(newNodes, node.Right)
		} else {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
			buffer.WriteString(" ")
		}
		buffer.WriteString(whitespaceN(space))
	}

	buffer.WriteString("\n")

	for i := 1; i <= edges; i++ {
		for _, node := range nodes {
			buffer.WriteString(whitespaceN(heads - i))
			if node == nil {
				buffer.WriteString(whitespaceN(2*edges + i + 1))
				continue
			}

			if node.Left != nil {
				buffer.WriteString("/")
			} else {
				buffer.WriteString(whitespaceN(1))
			}
			buffer.WriteString(whitespaceN(2*i - 1))

			if node.Right != nil {
				buffer.WriteString("\\")
			} else {
				buffer.WriteString(whitespaceN(1))
			}
			buffer.WriteString(whitespaceN(2*edges - i))

		}
		buffer.WriteString("\n")
	}

	vertical(newNodes, level+1, depth).WriteTo(buffer)

	return
}

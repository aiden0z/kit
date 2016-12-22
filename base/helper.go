package base

import "fmt"

type Int int

func (i Int) CompareTo(o Comparable) int {

	other, ok := o.(Int)
	if !ok {
		return 1
	}

	if i > other {
		return 1
	} else if i == other {
		return 0
	} else {
		return -1
	}
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

func NewIntComparableSlice(slice []int) (s []Comparable) {
	for _, v := range slice {
		s = append(s, Int(v))
	}
	return
}

type Rune rune

func (r Rune) CompareTo(o Comparable) int {
	other, ok := o.(Rune)
	if !ok {
		return 1

	}
	if r > other {
		return 1
	} else if r == other {
		return 0
	} else {
		return -1
	}
}

func (r Rune) String() string {
	return string(r)
}

func NewRuneComparableSlice(slice []rune) (r []Comparable) {
	for _, v := range slice {
		r = append(r, Rune(v))
	}

	return
}

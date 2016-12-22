package base

// Comparable describe a comparable interface
type Comparable interface {
	// CompareTo return a negative integer, zero, or a positive integer
	// as this object is less than, equal to, or greater than the specified object.
	CompareTo(o Comparable) int
	String() string
}

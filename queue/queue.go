// code from https://github.com/oleiade/lane, but something changes

package queue

// Queue is a FIFO (First in first out) data structure implementation.
// It is based on a deque container and focuses its API on core.
//
// As it is implemented using a Deque container, every operations
// over an Queue are synchronized and safe for concurrent
// usage.
//     <-- [ queue ] <--

type Queue struct {
	*Deque
}

func NewQueue() *Queue {
	return &Queue{
		Deque: NewDeque(),
	}
}

// Enqueue adds an item at the back of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.Append(item)
}

// Dequeue removes and returns the front queue item
func (q *Queue) Dequeue() interface{} {
	return q.Shift()
}

// Head returns the front queue item
func (q *Queue) Head() interface{} {
	return q.First()
}

package queue

import (
	"strconv"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	queue := NewQueue()
	queueSize := 100

	// Populate test queue and assert Enqueue
	// function does not fail
	for i := 0; i < queueSize; i++ {
		value := strconv.Itoa(i)
		queue.Enqueue(value)
	}

	if queue.Size() != queueSize {
		t.Error("queue size error")

	}

	if queue.First() != "0" {
		t.Error("queue first value error")

	}

	if queue.Last() != "99" {
		t.Error("queue first value error")

	}
}

func TestQueueDequeue_fulfilled(t *testing.T) {
	queue := NewQueue()
	queueSize := 100

	// Populate test queue and assert Enqueue
	// function does not fail
	for i := 0; i < queueSize; i++ {
		value := strconv.Itoa(i)
		queue.Enqueue(value)
	}

	// Check that while deuqueing, elements come out in
	// their insertion order
	for i := 0; i < queueSize; i++ {
		item := queue.Dequeue()
		expectedValue := strconv.Itoa(i)

		if item != expectedValue {
			t.Error(" queue dequue error")

		}

		if queue.Size() != queueSize-(i+1) {
			t.Error("queue size error")
		}
	}
}

func TestQueueDequeue_empty(t *testing.T) {
	queue := NewQueue()
	item := queue.Dequeue()

	if item != nil {
		t.Error("empty queue deque error")

	}

	if queue.Size() != 0 {
		t.Error("empty queue size error")
	}
}

func TestQueueHead_fulfilled(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("1")
	item := queue.Head()

	if item != "1" {
		t.Error("queue head error")

	}

	if queue.Size() != 1 {
		t.Error("queue size error")
	}
}

func TestQueueHead_empty(t *testing.T) {
	queue := NewQueue()
	item := queue.Head()

	if item != nil {
		t.Error("empty queue head error")

	}

	if queue.Size() != 0 {
		t.Error("empty queue size error")
	}
}

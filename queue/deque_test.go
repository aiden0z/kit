package queue

import (
	"strconv"
	"testing"
)

func TestDequeAppend(t *testing.T) {
	deque := NewDeque()
	sampleSize := 100

	// Append elements in the Deque and assert it does not fail
	for i := 0; i < sampleSize; i++ {
		value := strconv.Itoa(i)

		if err := deque.Append(value); err != nil {

			t.Error("Deque Append error")
		}
	}

	if deque.Size() != sampleSize {
		t.Error("deque size error")

	}

	if deque.First() != "0" {
		t.Error("deque first value error")

	}

	if deque.Last() != "99" {
		t.Error("deque last value error")

	}
}

func TestDequeAppendWithCapacity(t *testing.T) {
	dequeSize := 20
	deque := NewCappedDeque(dequeSize)

	// Append the maximum number of elements in the Deque
	// and assert it does not fail
	for i := 0; i < dequeSize; i++ {
		var value string = strconv.Itoa(i)

		if err := deque.Append(value); err != nil {

			t.Error("Deque Append error")
		}
	}

	// Try to overflow the Deque size limit, and make
	// sure appending fails
	if err := deque.Append("should not be ok"); err == nil {
		t.Error("deque should raise capacity full error")

	}
}

func TestDequePrepend(t *testing.T) {
	deque := NewDeque()
	sampleSize := 100

	// Prepend elements in the Deque and assert it does not fail
	for i := 0; i < sampleSize; i++ {
		value := strconv.Itoa(i)
		if err := deque.Prepend(value); err != nil {
			t.Error("deque prepend error")

		}

	}

	if deque.Size() != sampleSize {
		t.Error("deque size error")

	}

	if deque.First() != "99" {
		t.Error("deque first value error")

	}

	if deque.Last() != "0" {
		t.Error("deque last value error")

	}
}

func TestDequePrependWithCapacity(t *testing.T) {
	dequeSize := 20
	deque := NewCappedDeque(dequeSize)

	// Prepend elements in the Deque and assert it does not fail
	for i := 0; i < dequeSize; i++ {
		value := strconv.Itoa(i)
		if err := deque.Prepend(value); err != nil {
			t.Error("deque prepend error")

		}
	}

	// Try to overflow the Deque size limit, and make
	// sure appending fails
	if err := deque.Prepend("should not be ok"); err == nil {
		t.Error("deque should raise capacity full error")

	}

}

func TestDequePop_fulfilled_container(t *testing.T) {
	deque := NewDeque()
	dequeSize := 100

	// Populate the test deque
	for i := 0; i < dequeSize; i++ {
		value := strconv.Itoa(i)
		deque.Append(value)
	}

	// Pop elements of the deque and assert elements come out
	// in order and container size is updated accordingly
	for i := dequeSize - 1; i >= 0; i-- {
		item := deque.Pop()

		var itemValue string = item.(string)
		var expectedValue string = strconv.Itoa(i)

		if itemValue != expectedValue {
			t.Error("deque pop value error")

		}

		if deque.Size() != i {
			t.Error("deque size error")

		}

	}
}

func TestDequePop_empty_container(t *testing.T) {
	deque := NewDeque()
	item := deque.Pop()

	if item != nil {
		t.Error("pop empty deque error")

	}

	if deque.Size() != 0 {
		t.Error("deque size error")

	}
}

func TestDequeShift_fulfilled_container(t *testing.T) {
	deque := NewDeque()
	dequeSize := 100

	// Populate the test deque
	for i := 0; i < dequeSize; i++ {
		value := strconv.Itoa(i)
		deque.Append(value)
	}

	// Pop elements of the deque and assert elements come out
	// in order and container size is updated accordingly
	for i := 0; i < dequeSize; i++ {
		item := deque.Shift()

		itemValue := item.(string)
		expectedValue := strconv.Itoa(i)

		if itemValue != expectedValue {
			t.Error("deque shift value error")

		}

		if deque.Size() != (dequeSize - (i + 1)) {
			t.Error("deque size error")

		}

	}
}

func TestDequeShift_empty_container(t *testing.T) {
	deque := NewDeque()

	item := deque.Shift()

	if item != nil {
		t.Error("shift empty deque error")

	}

	if deque.Size() != 0 {
		t.Error("deque size error")

	}
}

func TestDequeFirst_fulfilled_container(t *testing.T) {
	deque := NewDeque()
	deque.Append("1")
	item := deque.First()

	if item != "1" {
		t.Error("deque first value error")

	}

	if deque.Size() != 1 {
		t.Error("deque size error")

	}

}

func TestDequeLast_fulfilled_container(t *testing.T) {
	deque := NewDeque()

	deque.Append("1")
	deque.Append("2")
	deque.Append("3")

	item := deque.Last()

	if item != "3" {
		t.Error("deque first value error")

	}

	if deque.Size() != 3 {
		t.Error("deque size error")

	}
}

func TestDequeEmpty_fulfilled(t *testing.T) {
	deque := NewDeque()
	deque.Append("1")

	if deque.IsEmpty() == true {
		t.Error("deque IsEmpty return error")

	}

}

func TestDequeEmpty_empty_deque(t *testing.T) {
	deque := NewDeque()
	if deque.IsEmpty() == false {
		t.Error("deque IsEmpty return error")

	}
}

func TestDequeFull_fulfilled(t *testing.T) {
	deque := NewCappedDeque(3)

	deque.Append("1")
	deque.Append("2")
	deque.Append("3")

	if deque.IsFull() == false {
		t.Error("deque IsFull return error")

	}
}

func TestDequeFull_non_full_deque(t *testing.T) {
	deque := NewCappedDeque(3)
	deque.Append("1")

	if deque.IsFull() == true {
		t.Error("deque IsFull return error")

	}
}

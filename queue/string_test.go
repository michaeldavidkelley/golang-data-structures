package queue

import (
	"testing"
)

func Test_FIFO(t *testing.T) {
	t.Parallel()

	expected := "the string"

	q := NewStringQueue()

	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue(expected)

	actual := q.Dequeue()

	if expected != actual {
		t.Fatalf("Expected (%s), Got (%s)", expected, actual)
	}
}

func Test_Dequeue_Empty_Queue(t *testing.T) {
	t.Parallel()

	q := NewStringQueue()

	if q.Dequeue() != "" {
		t.Fatal("Expected empty queue to Dequeue empty string")
	}
}

func Test_Size(t *testing.T) {
	t.Parallel()

	q := NewStringQueue()
	q.Enqueue("a")
	q.Enqueue("a")
	q.Enqueue("a")
	q.Enqueue("a")

	if q.Size() != 4 {
		t.Fatalf("Expected Size to be 4, Got %d", q.Size())
	}

	q.Dequeue()

	if q.Size() != 3 {
		t.Fatalf("Expected Size to be 3, Got %d", q.Size())
	}
}

func Test_IsEmpty(t *testing.T) {
	t.Parallel()

	q := NewStringQueue()

	if !q.IsEmpty() {
		t.Fatal("Expected Queue to be Empty")
	}

	q.Enqueue("a")

	if q.IsEmpty() {
		t.Fatal("Expected Queue not to be empty")
	}

	q.Dequeue()

	if !q.IsEmpty() {
		t.Fatal("Expected Queue to be Empty")
	}
}

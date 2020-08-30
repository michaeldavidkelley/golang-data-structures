package stack

import (
	"testing"
)

func Test_LIFO(t *testing.T) {
	t.Parallel()

	expected := "the string"

	stack := NewStringStack()

	stack.Push(expected)
	stack.Push("a")
	stack.Push("a")

	actual := stack.Pop()

	if expected != actual {
		t.Fatalf("Expected (%s), Got (%s)", expected, actual)
	}
}

func Test_Pop_Empty_Stack(t *testing.T) {
	t.Parallel()

	stack := NewStringStack()

	if stack.Pop() != "" {
		t.Fatal("Expcted empty stack to pop empty string")
	}
}

func Test_Size(t *testing.T) {
	t.Parallel()

	stack := NewStringStack()

	stack.Push("a")
	stack.Push("a")
	stack.Push("a")
	stack.Push("a")

	if stack.Size() != 4 {
		t.Fatalf("Expected 4, Got %d", stack.Size())
	}
}

func Test_IsEmtpy(t *testing.T) {
	t.Parallel()

	stack := NewStringStack()

	if !stack.IsEmpty() {
		t.Fatal("Expected Stack to be empty")
	}

	stack.Push("a")

	if stack.IsEmpty() {
		t.Fatal("Expected Stack NOT to be empty")
	}

	stack.Pop()

	if !stack.IsEmpty() {
		t.Fatal("Expected Stack to be empty")
	}
}

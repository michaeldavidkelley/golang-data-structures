package bag

import (
	"testing"
)

func Test_Bag(t *testing.T) {
	t.Parallel()

	bag := NewStringBag()

	if bag.Size() != 0 {
		t.Fatal("Expected bag size to be 0")
	}

	if !bag.IsEmpty() {
		t.Fatal("Expcted bag to be empty")
	}

	bag.Add("a")
	bag.Add("a")
	bag.Add("a")

	if bag.Size() != 3 {
		t.Fatal("Expected bag size to be 3")
	}

	if bag.IsEmpty() {
		t.Fatal("Expcted bag NOT to be empty")
	}
}

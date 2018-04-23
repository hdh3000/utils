package checklist

import (
	"testing"
)

func TestList(t *testing.T) {
	l := NewList("Test", NewItem("a"), NewItem("c"), NewItem("b"))

	if len(l.Elements()) != 3 {
		t.Errorf("list should have three elements it has %d", len(l.Elements()))
	}

	l.Push(NewItem("a1"))

	if len(l.Elements()) != 4 {
		t.Errorf("list should have four elements it has %d", len(l.Elements()))
	}

	if l.Elements()[3].Title() != "a1" {
		t.Errorf("list has the wrong final element")
	}

	l.Combine(l)
	if len(l.Elements()) != 4 {
		// Shouldn't be adding duplicates
		t.Errorf("list should have 4 elements it has %d", len(l.Elements()))
	}

	l2 := NewList("Test2", NewItem("d"), NewItem("e"))
	l.Combine(l2)
	if len(l.Elements()) != 6 {
		t.Errorf("list should have 6 elements it has %d", len(l.Elements()))
	}

	if l.Elements()[4].Title() != "d" {
		t.Errorf("list item 4 should be d")
	}

	li := l.NewInstance("test-instance")
	if len(li.All()) != 6 {
		t.Errorf("list should have 6 elements it has %d", len(l.Elements()))
	}

	if len(li.UnChecked()) != 6 {
		t.Errorf("list should have 6 unchecked elements it has %d", len(l.Elements()))
	}

	li.UnChecked()[1].Check()

	if len(li.UnChecked()) != 5 {
		t.Errorf("list should have 5 unchecked elements it has %d", len(l.Elements()))
	}

	if len(li.Checked()) != 1 {
		t.Errorf("list should have 1 checked elements it has %d", len(l.Elements()))
	}

}

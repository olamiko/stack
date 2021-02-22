package stack

import (
	"testing"
)

var newStack = Stack{}

func TestSize(t *testing.T) {
	description := "Test stack size function"
	if newStack.Size() != 0 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 0, newStack.Size())
	}
}

func TestPeek(t *testing.T) {
	newStack.Push(1)
	description := "Test stack peek function"
	if newStack.Peek() != 1 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 1, newStack.Peek())
	}

}

func TestPush(t *testing.T) {
	newStack.Push(3)
	description := "Test stack push function"
	if newStack.Peek() != 3 || newStack.Size() != 2 {
		t.Fatalf("FAIL: %s\nExpected Value: %#v\nActual Value: %#v", description, 3, newStack.Peek())
		t.Fatalf("FAIL: %#v\nActual Value: %#v", 2, newStack.Size())
	}

}

func TestPop(t *testing.T) {
	newStack.Push(5)
	description := "Test stack pop function"
	poppedValue := newStack.Pop()
	if poppedValue != 5 || newStack.Size() != 2 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 5, poppedValue)
		t.Fatalf("FAIL: %#v\nActual Value: %#v", 2, newStack.Size())
	}

}

func TestIsEmpty(t *testing.T) {
	description := "Test stack isEmpty function"
	if newStack.IsEmpty() {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 0, newStack.IsEmpty())
	}
}

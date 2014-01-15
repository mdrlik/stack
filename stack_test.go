package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	if !stack.IsEmpty() {
		t.Fatal("expected IsEmpty: true, got false")
	}
	stack.Push(0)
	if stack.IsEmpty() {
		t.Fatal("expected IsEmpty: false, got true")
	}
	data := stack.Pop()
	if data != 0 {
		t.Fatalf("expected data: 0, got %v", data)
	}
	stack.Push(3)
	*stack.Top() += 2
	data = *stack.Top()
	if data != 5 {
		t.Fatalf("expected data: 5, got %v", data)
	}
	stack = NewStack()
	stack.Push(3)
	stack.Push(5)
	data1 := stack.Pop()
	data2 := stack.Pop()
	if data1 != 5 {
		t.Fatalf("expected data: 5, got %v", data1)
	}
	if data2 != 3 {
		t.Fatalf("expected data: 3, got %v", data2)
	}
}

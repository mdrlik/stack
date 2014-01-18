package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	_, err := stack.Pop()
	if err != ErrStackIsEmpty {
		t.Fatalf("expected error %v, got %v", ErrStackIsEmpty, err)
	}
	stack.Push(0)
	data, err := stack.Pop()
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if data != 0 {
		t.Fatalf("expected data 0, got %v", data)
	}
	_, err = stack.Pop()
	if err == nil || err != ErrStackIsEmpty {
		t.Fatalf("expected error %v, got %v", ErrStackIsEmpty, err)
	}
	stack.Push(5)
	stack.Push(6)
	data1, err := stack.Pop()
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if data1 != 6 {
		t.Fatalf("expected data1 6, got %v", data1)
	}
	data2, err := stack.Pop()
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if data2 != 5 {
		t.Fatalf("expected data2 5, got %v", data2)
	}
}

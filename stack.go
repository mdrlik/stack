// Package stack implements LIFO data structure.
package stack

import (
	"errors"
)

var (
	ErrStackIsEmpty = errors.New("stack: stack is empty")
)

type node struct {
	data int
	back *node
}

type Stack struct {
	node *node
}

func NewStack() *Stack {
	return &Stack{}
}

// Pop removes and returns the data at the top of the stack. Returns
// ErrStackIsEmpty if stack is empty.
func (stack *Stack) Pop() (int, error) {
	if stack.node == nil {
		return 0, ErrStackIsEmpty
	}
	data := stack.node.data
	stack.node = stack.node.back
	return data, nil
}

// Push pushes a data at the top of the stack.
func (stack *Stack) Push(data int) {
	stack.node = &node{data, stack.node}
}

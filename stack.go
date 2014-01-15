// Package stack implements LIFO data structure.
package stack

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

func (stack *Stack) IsEmpty() bool {
	return stack.node == nil
}

func (stack *Stack) Push(data int) {
	stack.node = &node{data, stack.node}
}

func (stack *Stack) Pop() int {
	data := stack.node.data
	stack.node = stack.node.back
	return data
}

func (stack *Stack) Top() *int {
	return &stack.node.data
}

package main

type Stack struct {
	list SinglyLinkedList
}

func (stack *Stack) Enqueue(value interface{}) {
	stack.list.InsertBefore(value)
}

func (stack *Stack) Dequeue() interface{} {
	return stack.list.RemoveFirst()
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func NewStack() Stack {
	return Stack{
		list: NewSinglyLinkedList(),
	}
}

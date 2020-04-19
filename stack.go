package main

type Stack struct {
	list SinglyLinkedList
}

func (stack *Stack) Insert(value interface{}) {
	stack.list.InsertBefore(value)
}

func (stack *Stack) Remove() interface{} {
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

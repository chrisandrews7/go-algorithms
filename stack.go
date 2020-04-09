package main

type Stack struct {
	list SinglyLinkedList
}

func (stack *Stack) Insert(node *Node) {
	stack.list.InsertBefore(node)
}

func (stack *Stack) Remove() *Node {
	return stack.list.RemoveFirst()
}

func NewStack() Stack {
	return Stack{
		list: NewSinglyLinkedList(),
	}
}

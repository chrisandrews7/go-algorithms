package main

type Queue struct {
	list SinglyLinkedList
}

func (queue *Queue) Insert(node *Node) {
	queue.list.InsertAfter(node)
}

func (queue *Queue) Remove() *Node {
	return queue.list.RemoveFirst()
}

func NewQueue() Queue {
	return Queue{
		list: NewSinglyLinkedList(),
	}
}

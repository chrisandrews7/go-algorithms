package main

type Queue struct {
	list SinglyLinkedList
}

func (queue *Queue) Insert(value interface{}) {
	queue.list.InsertAfter(value)
}

func (queue *Queue) Remove() interface{} {
	return queue.list.RemoveFirst()
}

func NewQueue() Queue {
	return Queue{
		list: NewSinglyLinkedList(),
	}
}

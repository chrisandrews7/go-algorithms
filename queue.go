package main

type Queue struct {
	list SinglyLinkedList
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.list.InsertAfter(value)
}

func (queue *Queue) Dequeue() interface{} {
	return queue.list.RemoveFirst()
}

func (queue *Queue) Len() int {
	return queue.list.Len()
}

func NewQueue() Queue {
	return Queue{
		list: NewSinglyLinkedList(),
	}
}

package main

import (
	"errors"
)

type Node struct {
	value interface{}
	next  *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *SinglyLinkedList) InsertBefore(value interface{}) {
	node := &Node{value: value}

	list.length += 1

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	node.next = list.head
	list.head = node
}

func (list *SinglyLinkedList) InsertAfter(value interface{}) {
	node := &Node{value: value}

	list.length += 1

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	list.tail = node
}

func (list *SinglyLinkedList) RemoveFirst() interface{} {
	poppedNode := list.head

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		list.length = 0
		return poppedNode.value
	}

	list.head = list.head.next
	poppedNode.next = nil
	list.length -= 1

	return poppedNode.value
}

func (list *SinglyLinkedList) RemoveLast() interface{} {
	poppedNode := list.tail

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		list.length = 0
		return poppedNode.value
	}

	currentNode := list.head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}

	list.length -= 1
	currentNode.next = nil
	list.tail = currentNode

	return poppedNode.value
}

func (list *SinglyLinkedList) Remove(value interface{}) {
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		list.length = 0
		return
	}

	currentNode := list.head
	for currentNode.next.value != value && currentNode.next != nil {
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		panic(errors.New("Node not found"))
	}

	list.length -= 1
	currentNode.next = currentNode.next.next
	if currentNode.next == nil {
		list.tail = currentNode
	}
}

func (list *SinglyLinkedList) Len() int {
	return list.length
}

func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

package main

import (
	"errors"
)

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *SinglyLinkedList) InsertBefore(node *Node) {
	list.length += 1

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	node.next = list.head
	list.head = node
}

func (list *SinglyLinkedList) InsertAfter(node *Node) {
	list.length += 1

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	list.tail = node
}

func (list *SinglyLinkedList) RemoveFirst() *Node {
	poppedNode := list.head

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		list.length = 0
		return poppedNode
	}

	list.head = list.head.next
	poppedNode.next = nil
	list.length -= 1

	return poppedNode
}

func (list *SinglyLinkedList) RemoveLast() *Node {
	poppedNode := list.tail

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		list.length = 0
		return poppedNode
	}

	currentNode := list.head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}

	list.length -= 1
	currentNode.next = nil
	list.tail = currentNode

	return poppedNode
}

func (list *SinglyLinkedList) Remove(node *Node) {
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		list.length = 0
		return
	}

	currentNode := list.head
	for currentNode.next != node && currentNode.next != nil {
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

func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

func NewNode(value int) Node {
	return Node{
		value: value,
	}
}

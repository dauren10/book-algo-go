package main

import "errors"

type Node[T any] struct {
	data    T
	nextPtr *Node[T]
}

type SingleLinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func newSingleLinkedList[T any]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (sl *SingleLinkedList[T]) Size() int {
	return sl.length
}

func (sl *SingleLinkedList[T]) isEmpty() bool {
	return sl.length == 0
}

func newNode[T any](data T) *Node[T] {
	return &Node[T]{data, nil}
}

func (sl *SingleLinkedList[T]) pushTail(data T) error {
	node := newNode[T](data)
	if sl.length <= 0 {
		sl.head = node
		sl.tail = node
		sl.length = 1
		return nil
	}

	if sl.tail.nextPtr != nil {
		return errors.New("it is not a tail")
	}
	return nil
}

func (sl *SingleLinkedList[T]) pushHead(data T) error {
	return nil
}

package linkedlist

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

func newNode(value int) node {
	return node{
		value: value,
		next:  nil,
	}
}

type LinkedList struct {
	head, tail *node
	length     int
}

func NewLinkedList(value int) LinkedList {
	newNode := newNode(value)
	return LinkedList{
		head:   &newNode,
		tail:   &newNode,
		length: 1,
	}
}

func (l *LinkedList) PrintList() {
	var curr *node = l.head
	builder := strings.Builder{}
	for curr != nil {
		builder.WriteString(fmt.Sprint(curr.value) + " -> ")
		curr = curr.next
	}
	builder.WriteString(" nil ")
	fmt.Println(builder.String())
}

func (l *LinkedList) GetHead() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("list is empty")
	}
	return l.head.value, nil
}

func (l *LinkedList) GetTail() (int, error) {
	if l.tail == nil {
		return 0, fmt.Errorf("list is empty")
	}
	return l.tail.value, nil
}

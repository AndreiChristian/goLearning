package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Append(value int) {

	newNode := &Node{value, nil}

	if list.head == nil {

		list.head = newNode

	} else {

		lastNode := list.head
		for lastNode.next != nil {

			lastNode = lastNode.next

		}

		lastNode.next = newNode

	}

}

func (list *LinkedList) Prepend(value int) {

	newNode := &Node{value, nil}

	if list.head == nil {
		list.head = newNode
	} else {

		newNode.next = list.head
		list.head = newNode

	}

}

func (list *LinkedList) Display() {

	node := list.head

	for node != nil {
		fmt.Print(node.value, " ")
		node = node.next
	}

}

func (list *LinkedList) ToSlice() []int {

	if list.head == nil {
		return []int{}
	}

	node := list.head

	var slice []int

	for node != nil {
		slice = append(slice, node.value)
		node = node.next
	}

	return slice

}

func (list *LinkedList) isEmpty() bool {

	if list.head == nil {
		return true
	} else {
		return false
	}

}

func (list *LinkedList) length() int {

	if list.head == nil {
		return 0
	}

	length := 0

	for node := list.head; node != nil; {

		length++
		node = node.next

	}

	return length

}

func main() {

	l := LinkedList{}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	l.Prepend(5)
	l.Prepend(6)

	fmt.Println(l)

}

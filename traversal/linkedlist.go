package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepand(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	myList := linkedList{}

	node1 := &node{data: 10}
	node2 := &node{data: 30}

	myList.prepand(node1)
	myList.prepand(node2)

	fmt.Println(myList)
}

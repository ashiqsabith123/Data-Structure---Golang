package main

import "fmt"

type node struct {
	data int
	next *node
}

func createNode(value int) node {
	newNode := new(node)
	newNode.data = value

	return *newNode

}

type list struct {
	front *node
	back  *node
}

func (l *list) enQueue(element int) {

	node := createNode(element)

	if l.back == nil {
		l.front, l.back = &node, &node
	}

	l.back.next = &node
	l.back = &node
}

func (l *list) deQueue() int {
	var temp int
	if l.front == nil {
		fmt.Println("Queue is empty.....!")

	} else {
		temp = l.front.data
		l.front = l.front.next
	}

	return temp
}

func (d *list) display() {
	temp := d.front

	for temp != nil {
		fmt.Println(temp.data)

		temp = temp.next
	}
}

func main() {
	myList := list{}

	arr := []int{2, 2, 1}

	for i := 0; i < len(arr); i++ {
		myList.enQueue(arr[i])
	}

}

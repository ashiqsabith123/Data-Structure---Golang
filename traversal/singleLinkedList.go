package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type list struct {
	head *Node
	len  int
}

func createNode(data int) Node {
	newNode := new(Node)
	newNode.data = data

	return *newNode
}

func (l *list) search(value int) {
	temp := l.head

	for temp.next != nil {
		if temp.data == value {
			fmt.Println("The element found....")
			return
		}

		temp = temp.next
	}

	fmt.Println("The element not found....")
}

func (n *list) addNode(data int) {

	node := createNode(data)

	if n.head == nil {
		n.head = &node
		n.len++
	} else {

		ptr := n.head
		for i := 0; i < n.len; i++ {
			if ptr.next == nil {
				ptr.next = &node
				n.len++
				return

			} else {
				ptr = ptr.next

			}
		}

	}

}

func (d *list) insertBegin(value int) {
	node := createNode(value)

	second := d.head
	d.head = &node
	d.head.next = second
	d.len++

}

func (d *list) convertArrayToList(arr [10]int) {
	i := 0
	for i < len(arr) {
		d.addNode(arr[i])
		i++
	}
}

func (d *list) convertArrayToListRecur(arr [10]int, i int) {
	if i <= len(arr) {
		return
	}
	d.addNode(arr[i])

	d.convertArrayToListRecur(arr, i+1)
}

func (t *list) insertAfter(pos int, value int) {
	node := createNode(value)

	temp := t.head
	for temp != nil {

		if temp.data == pos {
			prev := temp.next
			temp.next = &node
			node.next = prev
			t.len++
			return

		}
		temp = temp.next
	}
	fmt.Println("Element not found........!")
}

func (d list) display() {

	if d.head == nil {
		fmt.Println("List is empty...!")
	} else {
		fmt.Print("The elements are: ")
		for i := d.head; i != nil; i = i.next {
			fmt.Printf("%d ", i.data)
		}
		fmt.Println("")
	}
}

func (r *list) replace(pos int, value int) {
	temp := r.head
	for temp != nil {
		if temp.data == pos {
			temp.data = value
			return
		}
		temp = temp.next
	}
	fmt.Println("Element not found")
}

func (s *list) Reverse() {

	current := s.head
	var before *Node
	var after *Node

	for current != nil {

		after = current.next
		current.next = before
		before = current
		current = after

	}
	s.head = before
}

func (g *list) reverseUsingRecursion(before *Node, after *Node, temp *Node) {

	after = temp.next
	temp.next = before
	before = temp
	temp = after

	if temp.next == nil {
		return
	}
	g.head = before

	g.reverseUsingRecursion(before, after, temp)

}

func (l *list) sort() {
	temp := l.head

	for temp != nil {

		temp1 := temp.next

		for temp1 != nil {

			if temp.data < temp1.data {

				prev := temp.data
				temp.data = temp1.data
				temp1.data = prev
			}

			temp1 = temp1.next

		}

		temp = temp.next

	}
}

func (l *list) delete(element int) {

	temp := l.head
	if temp != nil && temp.data == element {
		l.head = temp.next
		l.len--
		fmt.Println(element, "Deleted Succesfully...!")
		return
	}

	for temp.next != nil {

		if temp.next.data == element {

			temp.next = temp.next.next

			l.len--

			fmt.Println(element, "Deleted Succesfully...!")
			return
		}
		temp = temp.next

	}
	fmt.Println("Element not found......!")
}

func (e *list) deleteDuplicate() {
	temp := e.head

	for temp.next != nil {
		if temp.data == temp.next.data {
			temp.next = temp.next.next
			e.len--

		} else {
			temp = temp.next
		}

	}
}

func main() {
	myList := list{}

	arr := [10]int{1, 23, 9, 8, 89, 56, 13, 38, 80, 90}

	//myList.convertArrayToList(arr)
	var i int = 0
	myList.convertArrayToListRecur(arr, i)

	myList.addNode(44)
	myList.addNode(56)
	myList.addNode(99)
	myList.addNode(44)
	myList.addNode(86)
	myList.addNode(99)
	myList.addNode(98)
	myList.addNode(98)
	myList.insertBegin(3)
	myList.insertBegin(55)
	myList.insertAfter(99, 78)
	myList.insertAfter(78, 86)
	myList.insertAfter(56, 90)

	myList.delete(99)

	//fmt.Println(myList.head.next.next)
	myList.replace(98, 89)
	myList.addNode(98)
	myList.replace(98, 89)
	fmt.Println(myList.head.next.next.data)

	myList.replace(56, 78)
	myList.replace(44, 11)

	myList.search(44)
	myList.sort()

	myList.display()

	myList.deleteDuplicate()

	myList.Reverse()

	var before *Node
	var after *Node
	var temp *Node = myList.head

	myList.reverseUsingRecursion(before, after, temp)

	myList.display()

}

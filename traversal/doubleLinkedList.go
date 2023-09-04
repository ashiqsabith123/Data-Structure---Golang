package main

import "fmt"

type DNode struct {
	prev *DNode
	data int
	next *DNode
}

type DList struct {
	head *DNode
	tail *DNode
	len  int
}

func createDNode(value int) DNode {
	Dnode := new(DNode)
	Dnode.data = value

	return *Dnode
}

func (D *DList) addDNode(num int) {

	node := createDNode(num)

	if D.head == nil {
		D.head = &node
		D.len++
		return
	} else {
		temp := D.head
		for temp != nil {
			if temp.next == nil {
				temp.next = &node
				temp.next.prev = temp
				D.len++
				return

			}
			temp = temp.next
			D.tail = &node
		}

	}

	D.tail = &node
}

func (D *DList) insertAfter(pos int, element int) {
	node := createDNode(element)

	temp := D.head

	for temp != nil {
		if temp.data == pos {
			dat := temp.next
			temp.next = &node
			temp.next.prev = temp
			node.next = dat
			node.next.prev = &node
			return

		}
		temp = temp.next
	}

}

func (D *DList) insertBefore(pos int, element int) {
	node := createDNode(element)

	temp := D.head

	if D.head.data == pos {
		prev := D.head
		D.head = &node
		D.head.next = prev
		D.head.next.prev = &node
		return

	}

	for temp != nil {
		if temp.next.data == pos {
			dat := temp.next
			temp.next = &node
			temp.next.prev = temp
			node.next = dat
			node.next.prev = &node
			return

		}
		temp = temp.next
	}

}

func (D DList) display() {
	temp := D.head
	if D.head == nil {
		fmt.Printf("List is empty...")
	} else {
		fmt.Printf("The elements are: ")
		for temp != nil {
			fmt.Printf("%v ", temp.data)
			temp = temp.next
		}
		fmt.Println(" ")
	}

}

func (D *DList) delete(value int) {

	if D.head.data == value && D.head.next != nil {
		D.head = D.head.next
		D.head.prev = nil
		return
	} else if D.head.next == nil {
		D.head = nil
	} else {
		temp := D.head
		for temp != nil {
			if temp.next.data == value && temp.next.next != nil {
				pre := temp
				temp.next = temp.next.next
				temp.next.prev = pre
				return
			} else if temp.next.data == value {
				temp.next = nil
			}
			temp = temp.next
		}
	}
}

func (D DList) Rdisplay() {
	temp := D.tail
	fmt.Printf("The elements are in reverse order: ")
	for temp != nil {
		fmt.Printf("%v ", temp.data)
		temp = temp.prev
	}
	fmt.Println(" ")
}

func (D *DList) center() {
	i := 1
	temp := D.head
	val := D.len / 2

	if val%2 == 1 {
		val = val + 1
	}
	for i <= val {
		if i == val {
			fmt.Println(temp.data)
			return
		}
		i++
		temp = temp.next
	}
}

func main() {
	myDList := DList{}

	myDList.addDNode(22)
	myDList.addDNode(33)
	//myDList.addDNode(89)
	//myDList.addDNode(87)
	// myDList.addDNode(78)
	fmt.Println(myDList.len)
	myDList.center()

	// myDList.insertAfter(33, 22)
	// myDList.insertAfter(22, 18)

	// myDList.insertBefore(33, 99)
	// myDList.insertBefore(18, 98)
	// myDList.insertBefore(22, 99)
	// myDList.insertBefore(99, 56)
	// myDList.insertBefore(33, 00)
	// myDList.delete(56)
	// myDList.display()

	// fmt.Println("******************")
	// myDList.delete(98)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")
	// myDList.delete(18)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")
	// myDList.delete(99)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")
	// myDList.delete(0)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")
	// myDList.delete(33)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")
	// myDList.delete(89)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")
	// myDList.delete(99)
	// myDList.display()
	//myDList.Rdisplay()
	//fmt.Println("******************")

	myDList.display()

	// myDList.delete(22)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")

	// myDList.delete(87)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")

	// myDList.delete(22)
	// myDList.display()
	// //myDList.Rdisplay()
	// fmt.Println("******************")

	// fmt.Println(myDList.head.next.next.prev)

}

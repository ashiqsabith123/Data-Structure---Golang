package main

import "fmt"

type DONode struct {
	next *DONode
	data int
}

type DOList struct {
	head *DONode
	len  int
}

func (D *DOList) addNode(value int) {
	node := new(DONode)
	node.data = value

	if D.head == nil {
		D.head = node
		D.len++
		return
	} else {
		temp := D.head
		for temp != nil {
			if temp.next == nil {
				temp.next = node
				D.len++
				return
			}
			temp = temp.next
		}
	}
}

func (D *DOList) deleteDuplicate() {
	temp := D.head

	for temp != nil {
		ltemp := temp
		for ltemp != nil && ltemp.next != nil {
			if temp.data == ltemp.next.data {
				ltemp.next = ltemp.next.next
			} else {
				ltemp = ltemp.next
			}
		}

		temp = temp.next
	}
}

func (D *DOList) deleteGPTDuplicate() {
	temp := D.head

	m := make(map[int]bool)

	for temp != nil && temp.next != nil {
		if m[temp.next.data] {
			temp.next = temp.next.next
		} else {
			m[temp.data] = true
			temp = temp.next
		}
	}
}

func (D DOList) display() {
	temp := D.head

	for temp != nil {
		fmt.Println(temp)
		temp = temp.next
	}
}

func (D *DOList) deleteAlter() {
	temp := D.head

	for temp != nil && temp.next != nil {
		temp.next = temp.next.next

		temp = temp.next
	}
}

func (D *DOList) reverseUseIter() {
	temp := D.head

	

	D.head = temp

}

func main() {

	myList := DOList{}

	myList.addNode(1)
	myList.addNode(5)
	myList.addNode(1)
	myList.addNode(2)
	myList.addNode(2)
	myList.addNode(3)
	myList.addNode(4)

	myList.addNode(50)
	myList.addNode(50)
	myList.addNode(50)
	myList.addNode(60)
	myList.addNode(60)
	myList.addNode(70)
	myList.addNode(80)
	myList.addNode(80)
	myList.addNode(80)
	myList.addNode(80)

	myList.display()

	//myList.deleteAlter()
	//myList.deleteGPTDuplicate()

	myList.reverseUseIter()

	fmt.Println("---------")

	myList.display()

}

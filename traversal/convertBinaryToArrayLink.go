package main

import "fmt"

type CNode struct {
	next *CNode
	data int
}

type Clist struct {
	head *CNode
	len  int
}

func (C *Clist) addNodeFromArray(arr []int) {

	for _, v := range arr {
		node := new(CNode)
		node.data = v

		if C.head == nil {
			C.head = node
			C.len++
		} else {
			temp := C.head
			for temp != nil {
				if temp.next == nil {
					temp.next = node
					break
				} else {
					temp = temp.next
				}

			}
		}

	}

}

func (C *Clist) display() {
	temp := C.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main() {

	myList := Clist{}

	arr := []int{1, 2, 3, 4, 5, 6, 7}

	myList.addNodeFromArray(arr)

	myList.display()

}

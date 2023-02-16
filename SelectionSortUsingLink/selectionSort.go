package main

import "fmt"

type Dnode struct {
	prev *Dnode
	data int
	next *Dnode
}

type Dlist struct {
	head *Dnode
	tail *Dnode
}

func (n *Dlist) addNode(value int) {
	newNode := new(Dnode)
	newNode.data = value

	if n.head == nil {
		n.head = newNode
		return
	} else {
		temp := n.head

		for temp != nil {
			if temp.next == nil {
				temp.next = newNode
				newNode.prev = temp
				n.tail = newNode
				return

			}

			temp = temp.next
		}

	}
}

func (d *Dlist) selectionSort() {
	temp := d.head

	var tempNode *Dnode

	for temp != nil {
		small := temp.data
		temp1 := temp
		for temp1 != nil {
			//fmt.Println(tempNode)

			if temp1.data < small {
				tempNode = temp1
				small = temp1.data
			}
			temp1 = temp1.next
		}

		if tempNode != nil && tempNode.data < temp.data {
			temp.data, tempNode.data = small, temp.data
		}

		temp = temp.next
	}

}

func main() {
	list := Dlist{}

	list.addNode(1)
	list.addNode(89)
	list.addNode(99)
	list.addNode(78)
	list.addNode(38)
	list.addNode(0)
	list.addNode(11)

	t := list.head

	for t != nil {
		fmt.Println(t.data)

		t = t.next
	}

	list.selectionSort()

	m := list.head

	for m != nil {
		fmt.Println(m.data)

		m = m.next
	}

}

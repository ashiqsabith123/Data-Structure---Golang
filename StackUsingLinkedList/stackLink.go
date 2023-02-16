package main

import "fmt"

type data struct {
	element int
	next    *data
}

type stack struct {
	top *data
}

func (s *stack) push(value int) {
	newNode := new(data)
	newNode.element = value

	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}

}

func (s *stack) pop() {
	s.top = s.top.next
}

func (s *stack) display() {
	temp := s.top

	for temp != nil {
		fmt.Println(temp.element)

		temp = temp.next
	}
}

func main() {
	myStack := stack{}

	myStack.push(10)
	myStack.push(56)
	myStack.push(667)
	myStack.push(99)

	myStack.display()
	myStack.pop()
	myStack.pop()
	fmt.Println("-----------")
	myStack.display()

}

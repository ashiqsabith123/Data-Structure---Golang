package main

import "fmt"

const size = 10

type stack struct {
	data [size]int
	top  int
}

func (s *stack) push(value int) {
	if s.top == size-1 {
		fmt.Println("Stack overflow")
		return
	}
	s.top++
	s.data[s.top] = value
}

func (s *stack) pop() int {
	if s.top == -1 {
		fmt.Println("Stack underflow")
		return 0
	}
	value := s.data[s.top]
	s.top--
	return value
}

func main() {
	var s stack
	s.top = -1
	s.push(1)
	s.push(2)
	//s.push(3)
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s)
}

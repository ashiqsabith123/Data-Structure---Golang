package main

import "fmt"

func main() {

	array := []int{5, 6, 4, 1, 6, 8, 7, 6}

	num := enterTheValue()

	fmt.Println(delete(num, array))
}

func enterTheValue() int {
	var num int

	fmt.Printf("Enter The Element You Want To Delete ")
	fmt.Scan(&num)

	return num
}

func delete(num int, array []int) []int {

	var stat bool
	for i, d := range array {

		if d == num {
			array = append(array[:i], array[i+1:]...)

		} else {
			stat = true

		}
	}

	return array

	if stat {
		fmt.Println("Enter a valid element")
	}
	//size--

	return nil
}

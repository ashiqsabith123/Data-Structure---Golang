package main

import "fmt"

func () {

	array := [8]int{6, 5, 6, 6, 5, 4, 6, 6}

	number := read()

	fmt.Println(sort(number, array))

}

func read() int {
	var num int
	fmt.Printf("Enter the number you need to sort: ")
	fmt.Scan(&num)

	return num
}

func sort(value int, arr [8]int) [8]int {

	rearPos := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		if i < rearPos {
			for arr[rearPos] == value {
				rearPos--

			}

			if arr[i] == value {
				temp := arr[rearPos]
				arr[rearPos] = arr[i]
				arr[i] = temp
				rearPos--
			}
		}
	}

	return arr

}

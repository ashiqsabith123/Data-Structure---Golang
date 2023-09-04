package main

import "fmt"

func main() {

	numbers := [5]int{1, 2, 3, 4, 5}
	var temp int
	large := numbers[0]

	for _, i := range numbers {

		if i > large {
			temp = large
			large = i
		}

	}

	fmt.Println(temp)
}

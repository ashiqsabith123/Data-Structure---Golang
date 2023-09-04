package main

import "fmt"

func main() {
	arr := []int{1, -1, 2, 4, 6, -7, 8, 9}

	large := arr[0]
	small := arr[0]

	for _, v := range arr {
		if v > large {
			large = v
		}
	}

	for _, i := range arr {
		if i < small {
			small = i
		}
	}

	fmt.Println("small:", small)

	fmt.Println("large:", large)
}

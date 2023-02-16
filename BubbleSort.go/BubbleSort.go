package main

import "fmt"

func bubbleSort(array [10]int) {

	size := len(array) - 1

	for i := 0; i < len(array); i++ {
		for j := 0; j < size; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}

		size--
	}

	fmt.Println(array)
}

func main() {

	array := [10]int{23, 1, 45, 67, 78, 2, 90, 45, 2, 100}

	bubbleSort(array)

}

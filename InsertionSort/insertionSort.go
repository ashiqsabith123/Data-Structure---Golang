package main

import "fmt"

func sort(array [10]int) {
	for i := 1; i < len(array); i++ {
		current := array[i]

		j := i - 1

		for j >= 0 && array[j] > current {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = current
	}

	fmt.Println(array)
}

func main() {

	arr := [10]int{33, 2, 43, 1, 56, 78, 9, 0, 34, 56}

	sort(arr)

}

package main

import "fmt"

func sort(array [10]int) {

	for i := 0; i < len(array); i++ {
		small := array[i]
		var temp int
		for j := i + 1; j < len(array); j++ {
			if array[j] < small {
				temp = j
				small = array[j]
			}

		}

		if small < array[i] {
			array[i], array[temp] = small, array[i]
		}

	}

	fmt.Println(array)

}

func main() {
	arr := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	sort(arr)
}

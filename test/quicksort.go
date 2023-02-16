package main

import "fmt"

func quickSort(array []int, lower int, upper int) {

	if lower < upper {
		pivot := partition(array, lower, upper)
		quickSort(array, lower, pivot-1)
		quickSort(array, pivot+1, upper)
	}

}

func partition(array []int, lower int, upper int) int {

	pivot := array[upper]

	slow := lower - 1

	for i := lower; i <= upper; i++ {
		if array[i] < pivot {
			slow++
			array[slow], array[i] = array[i], array[slow]
		}
	}

	array[slow+1], array[upper] = pivot, array[slow+1]

	return slow + 1

}

func main() {

	arr := []int{2, 33, 1, 54, 67, 89, 52, 34, 19, 0, -3}
	middle := len(arr)
	quickSort(arr, 0, middle-1)
	fmt.Println(arr)

}

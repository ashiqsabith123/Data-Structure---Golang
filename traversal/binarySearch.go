package main

import "fmt"

func binarySearch(arr [5]int, target int) int {
	findex := 0
	lindex := len(arr) - 1

	//middle := lindex - findex

	for findex <= lindex {

		middle := (lindex + findex) / 2

		if arr[middle] == target {
			fmt.Println("found")
			return middle
		} else if target > arr[middle] {
			findex = middle + 1
		} else {
			lindex = middle - 1
		}

	}
	return 0
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(arr, 3))
}

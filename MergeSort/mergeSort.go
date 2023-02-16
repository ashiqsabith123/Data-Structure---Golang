package main

import "fmt"

func merge(left, right []int) []int {
	i, j := 0, 0
	var tempArray []int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tempArray = append(tempArray, left[i])
			i++
		} else {
			tempArray = append(tempArray, right[j])
			j++
		}
	}

	tempArray = append(tempArray, left[i:]...)
	tempArray = append(tempArray, right[j:]...)

	return tempArray
}

func mergeSort(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2
	left := mergeSort(array[:middle])
	right := mergeSort(array[middle:])
	return merge(left, right)

}

func main() {

	arr := []int{22, 1, 33, 2, 5, 67, 87, -9, 0, -1, 45}

	sorted := mergeSort(arr)
	fmt.Println(sorted)

}

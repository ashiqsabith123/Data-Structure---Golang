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

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])
	return merge(left, right)

}

func main() {
	arr := []int{22, 1, 33, 2, 5, 67, 87, -9, 0, -1, 45}
	fmt.Println(mergeSort(arr))
}

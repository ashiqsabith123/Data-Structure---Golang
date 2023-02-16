package main

import (
	"fmt"
)

func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func mergeSort(numbers []int) []int {

	if len(numbers) <= 1 {
		return numbers
	}
	middle := len(numbers) / 2
	left := mergeSort(numbers[:middle])
	right := mergeSort(numbers[middle:])
	return merge(left, right)
}

func main() {
	numbers := []int{5, 4, 1, 8, 7, 2, 6, 3, 9}
	fmt.Println("Before sorting: ", numbers)
	sortedNumbers := mergeSort(numbers)
	fmt.Println("After sorting: ", sortedNumbers)
}

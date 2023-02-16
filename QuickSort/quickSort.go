package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := Partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func Partition(array []int, lower int, upper int) int {
	pivot := array[upper]

	slow := lower - 1

	for j := lower; j < upper; j++ {
		if array[j] < pivot {
			slow++
			array[slow], array[j] = array[j], array[slow]
		}
	}

	array[slow+1], array[upper] = array[upper], array[slow+1]

	return slow + 1

}

func main() {
	Arr := []int{33, 4, 3, 2, 88, 78, 65, 43, -1, 56, 55}
	N := len(Arr)
	quickSort(Arr, 0, N-1)
	fmt.Println("Sorted array is:", Arr)

}

package main

import "fmt"

func insert(arr *[]int, pos int, value int) {
	*arr = append(*arr, 0)

	for i := len((*arr)) - 1; i >= pos; i-- {
		(*arr)[i] = (*arr)[i-1]
	}

	(*arr)[pos] = value

	fmt.Println(arr)
}

func aDelete(arr []int) {

	var num int

	fmt.Printf("The array elements are: ")
	fmt.Println(arr)

	fmt.Printf("Enter the element you want to delete: ")
	fmt.Scanf("%v", &num)

	for ind, v := range arr {
		if v == num {
			arr = append(arr[:ind], arr[ind+1:]...)
			break
		}
	}

	fmt.Println("After deleting ", num)
	fmt.Println(arr)

}

func deleteDuplicateArray(arr[] int){
	for i:=0; i< len(arr); i++{
		
	}
}

func insertAfterElement(arr []int, element int, insert int) {
	for ind, v := range arr {
		if v == element {
			arr = append(arr, 0)
			for i := len(arr) - 1; i >= ind+1; i-- {
				arr[i] = arr[i-1]
			}
			arr[ind+1] = insert
			break
		}
	}

	fmt.Println(arr)

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	//aDelete(arr)

	// insert(&arr, 4, 3)

	insertAfterElement(arr, 5, 10)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	// array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// target := 10

	// traverse(array, target)

	// array1 := [8]int{1, 2, 3, 4, 5, 6}
	// target1 := 10

	// value := traverse(array1, target1)

	// fmt.Println(value)

}

// func traverse(arr [8]int, target int) [2]int {

// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j < len(arr); j++ { // time complexity O(n2)
// 			if arr[i]+arr[j] == target {
// 				value := [2]int{arr[i], arr[j]} // space complexity 0(1)
// 				return value
// 			}
// 		}
// 	}

// 	ee := [2]int{2, 3}

// 	return ee
// }

func traversesp(ar [8]int, target int) []int {
	var value = make([]int, 0)

	//fmt.Println(value)
	for ; target < target; target++ {

	}

	for i := 0; i < len(ar); i++ {
		num := ar[i]
		match := target - num

		fmt.Println("ma", match)

		o := sort.SearchInts(value, match)

		fmt.Println("", o)

		if i != o {
			var val []int
			//fmt.Println("entered")
			val = append(val, num, match)
			return val

		} else {
			value = append(value, num)
			fmt.Println("value", value[i])
		}
	}

	ee := []int{2, 3}

	return ee
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

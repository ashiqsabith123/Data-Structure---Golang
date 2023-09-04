package main

import (
	"fmt"
	"math"
)

var sum int

var j int = 1
var k int = 1

func convert(array []int) {

	for i := len(array) - 1; i >= 0; i-- {
		// if i == len(array)-1 {
		// 	sum = array[i] * 1

		// } else {

		sum = sum + array[i]*(k*j)
		k = 2
		j = j * k

		// }

	}

	
	fmt.Println(sum/2 + 1)

}

func main() {
	temp := []int{1, 0, 1, 0, 1}

	convert(temp)
}

package main

import (
	"fmt"
	"strconv"
)

func stringChange(arr string, count int) {
	var temp []byte

	for i := 0; i < len(arr); i++ {
		position := arr[i] + byte(count)

		if position <= 122 {
			temp = append(temp, position)
		} else {
			position = position - 26
			temp = append(temp, byte(position))

		}

	}

	fmt.Println(string(temp))

}

func stringReverse() {
	tring := "Ashiq"

	var char []byte

	for i := len(tring) - 1; i >= 0; i-- {
		fmt.Println(string(tring[i]))

		char = append(char, tring[i])

	}

	fmt.Println(string(char))
}

func continuos(arr string) {
	count := 1
	var st string
	for i := 0; i <= len(arr); i++ {

		if i+1 >= len(arr) {
			break
		}

		if arr[i] == arr[i+1] {
			count = count + 1
		} else {
			fmt.Println(count)
			st = st + string(arr[i]) + strconv.Itoa(count) + ":"
			count = 1
		}

	}

	fmt.Println(st)
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	fmt.Println("doing")
	return i * factorial(i-1)
}

func stringE() {
	str := "Ashiq"
	str1 := "Sabith"

	str = str + str1

	fmt.Println(str)
}

func main() {
	//stringReverse()

	//arr := "HELLO"

	// f := "absd"
	//continuos(arr)
	//stringChange(arr, 2)
	// i := 5
	// fmt.Println(factorial(i))

	stringE()
}

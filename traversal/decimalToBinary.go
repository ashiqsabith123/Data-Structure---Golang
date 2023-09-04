package main

import (
	"fmt"
	"strconv"
)

func decimalToBinary(decimal int) string {
	binary := ""
	for decimal > 0 {
		binary = strconv.Itoa(decimal%2) + binary
		decimal = decimal / 2
	}
	return binary
}

func main() {
	decimal := 56
	binary := decimalToBinary(decimal)
	fmt.Println(binary)
}

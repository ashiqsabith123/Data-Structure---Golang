package main

import "fmt"

func main() {

	type sanika struct {
		name string
		age  int
	}

	p := sanika{
		name: "mendi",
		age:  88,
	}

	d := new(sanika)

	d.age = 12
	d.name = "rrr"

	fmt.Println(p)

	fmt.Println(d)

	fmt.Println(p)

}

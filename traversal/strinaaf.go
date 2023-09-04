package main

import "fmt"

func insertIntoString(str string, after string, word string) {

	var tt []byte
	var ttr string
	var j int = 0

	for i := 0; i <= len(str)-1; i++ {
		j = i
		for str[j] != 32 {
			//fmt.Println("doing")
			tt = append(tt, str[j])

			i = j + 1
			j++

			if j >= len(str) {
				break
			}
		}

		//fmt.Println(string(tt))

		if string(tt) != after {
			ttr = ttr + string(tt) + " "
		} else {
			ttr = ttr + string(tt) + " "
			ttr = ttr + word + " "
		}

		tt = nil

		//fmt.Println(i)

	}

	fmt.Println(ttr)

}

func main() {
	str := "my name is ashiq"

	insertIntoString(str, "name", "sss")

}

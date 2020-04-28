package main

import "fmt"

func main() {
	var number [7]int
	fmt.Println(number)
	y := 20
	for x := 2; x <= y; x++ {
		if y%2 == 0 {
			fmt.Println(x)
		}

	}
}

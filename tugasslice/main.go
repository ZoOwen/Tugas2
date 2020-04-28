package main

import "fmt"

func main() {

	s := make([]int, 0, 20)
	y := 20
	for x := 1; x < y; x++ {
		data := 0
		for a := 1; a < y; a++ {
			if x%a == 0 {
				data++
			}
		}
		if data == 2 && y > 1 {
			s = append(s, x)
		}

	}
	fmt.Println(s)

}

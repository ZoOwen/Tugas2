package main

import "fmt"

func main() {
	number := [30]int{1, 2, 3, 4, 5, 6, 7, 8, 19, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var ganjil []int
	var genap []int
	var prima []int

	fmt.Println("data s", number)
	for i := 0; i < len(number); i++ {
		if number[i]%2 == 0 {
			// fmt.Println("genap", number[i])
			genap = append(genap, number[i])
		} else {

			// fmt.Println("ganjil", number[i])
			ganjil = append(ganjil, number[i])
		}
	}

	for a := 1; a < len(number)-1; a++ {
		if number[a]%2 != 0 {
			// fmt.Println("prima", number[a])
			prima = append(prima, number[a])
		}
	}

	fmt.Println("total genap", genap, "jumlah data genap", len(genap))
	fmt.Println("total bilangan", ganjil, "jumlah data ganjil", len(ganjil))
	fmt.Println("total bilangan", prima, "jumlah data PRIMA", len(prima))

}

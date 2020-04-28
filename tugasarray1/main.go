package main

import (
	"fmt"
)

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func main() {
	array := [...]int{1, 7, 7, 19, 1, 18, 5, 0, 16, 0, 14, 11, 2, 9, 8, 14, 11, 5, 17, 6}
	array2 := [...]int{15, 6, 8, 18, 7, 77, 17, 8, 10, 15, 1, 8, 7, 11, 9, 16, 17, 11, 5, 6}

	slice1 := array[0:]
	slice2 := array2[0:]
	slice3 := append(slice1, slice2...)
	uniqueSlice := unique(slice3)
	fmt.Println("array1", array)
	fmt.Println("array2", array2)

	fmt.Println("angka yang tidak double", uniqueSlice)

}

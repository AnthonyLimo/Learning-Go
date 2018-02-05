//a slice is an array like data strucure without the constraints of how many
//values can be held in the array

package main

import "fmt"

func main() {
	favSlice := []int{5, 7, 7, 8, 9, 3, 2, 5, 6, 8, 5}

	favSlice2 := favSlice[4:7]

	fmt.Println("FavSlice -->")

	for _, value := range favSlice {
		fmt.Println(value)
	}

	fmt.Println("FavSlice2 -->")

	for _, value2 := range favSlice2 {
		fmt.Println(value2)
	}

}

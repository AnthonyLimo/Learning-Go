//VARIATIC FUNCTION
//This is a function that basically takes an unknown number of values and does
//a particular operation with that data

package main

import "fmt"

func main() {
	fmt.Println(subThem(1, 2, 3, 4, 5, 6, 7, 8))
}

func subThem(args ...int) int {
	result := 0

	for _, value := range args {
		fmt.Println(value)
		result -= value
	}

	return result
}

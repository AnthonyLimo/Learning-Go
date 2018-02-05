package main

import "fmt"

func main() {
	numSlice := []int{5, 4, 3, 2, 1}
	numSlice3 := make([]int, 5, 10)

	copy(numSlice3, numSlice)

	fmt.Println(numSlice3[0])

	numSlice3 = append(numSlice3, 0, 23)

	for _, value := range numSlice3 {
		fmt.Println(value)
	}
}

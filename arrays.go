package main

import "fmt"

func main() {
	favNums := [5]float64{4, 5, 6, 7, 8}
	fmt.Println("Index" + " Value")
	for i, value := range favNums {
		fmt.Println(i, value)
	}
}

//this is how you catch errors in your program without having the entire thing crash

package main

import "fmt"

func main(){
	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(434,5))
}

func safeDiv(num1, num2 int) int{
	defer func(){
		fmt.Println(recover())
	}()

	solution := num1/num2
	return solution
}

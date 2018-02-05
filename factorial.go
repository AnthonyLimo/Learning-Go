//this is a program calculating the factorial of a number
//it uses a recursive function which basically means a functions that calls itself
//... within itself

package main

import "fmt"

func main() {
	fmt.Println(factorial(34))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1) //function calls itself
}

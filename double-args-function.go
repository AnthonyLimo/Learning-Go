//allows a function to return two values from a function which is
//not supported in many programming languages

package main

import "fmt"

func main() {
	num1, num2 := nextTwoValues(7)
	fmt.Println(num1, num2)
}

func nextTwoValues(number int) (int, int) {
	return number + 1, number + 2
}

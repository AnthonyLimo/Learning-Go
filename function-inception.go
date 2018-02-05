//how to write a function inside another function
//this kinda looks like how you do it in javascript
//when placing functions inside objects

package main

import "fmt"

func main(){
	num := 3
	//declaring a function within a function
	doubleNum := func() int {
		num *= 2
		return num
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}

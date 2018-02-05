//defer function allows a linked function to run after the main function has run
//this means that it will delay the running of teh function until the whole thread is run

package main

import "fmt"

func main(){
	defer printTwo()
	printOne()
}

func printOne(){ fmt.Println(1) }
func printTwo() { fmt.Println(2) }

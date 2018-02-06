//how to raise an exception in go with the panic and recover function

package main

import "fmt"

func main(){
	demPanic()
}

func demPanic(){
	defer func(){
		fmt.Println(recover())
	}()

	panic("Panic!!!")
}

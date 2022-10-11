package main

import "fmt"

func main() {
	fmt.Println("Welcome to a call on Pointers ")

	/*var ptr *int
	fmt.Println("default value ", ptr)*/

	myNumber := 23
	var ptr = &myNumber

	//refrance means & sign
	//*ptr means the what inside in the ptr

	fmt.Println("value of pointer", ptr)
	fmt.Println("value of pointer", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new valuse is :", myNumber)

}

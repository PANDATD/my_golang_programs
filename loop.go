package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println("Hello world!", i)
	}
}

/*
	In Go programming languge their is only one loop
	such as for loop
	if we have to exceuter a block of code repetedly then we will use for loop.
	along with for keyword initalze varable put semicolon put condition again semicolon;
	and increment or decrement
	syntax:-

		for initalization; condition; increment/decrement{
		statement
		}

	Example:-

		for i := 1; i<=10; i++{
		fmt.Println("Hello World! ",i)
		}
*/

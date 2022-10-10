// This is exercise file for the golang
package main

import "fmt"

func main() {

	// Print Line in Go
	fmt.Println("Hello there, welcome to Golang From Pandatd")

	// Using Variables and Printing Theam

	var empName string = "Tejas Dixit"
	var empAge int64 = 21
	var empEmail string = "tejasdixit17@go.dev.com"
	var isMale bool = true

	fmt.Printf("Name :%v Type: %T \n", empName, empName)
	fmt.Printf("Age :%v Type: %T \n", empAge, empAge)
	fmt.Printf("Email :%v Type: %T \n", empEmail, empEmail)
	fmt.Printf("Male :%v Type: %T \n", isMale, isMale)

	// Now we will be Look for how arrays are created in golang how to access it ....

	var channels [4]string
	// keyWord var arrayName [sizeOfArray] typeOfArray

	channels[0] = "ch-00"
	channels[1] = "ch-01"
	channels[2] = "ch-02"
	channels[3] = "ch-03"

	fmt.Println(channels)
	fmt.Println("ch-01--->", channels[1])

	// Now we will look about How can create the Map in Golang map is nothig but Dictonery like in Python.

	numSpeel := make(map[int64]string)

	numSpeel[0] = "one"
	numSpeel[1] = "two"
	numSpeel[3] = "Three"

	fmt.Printf("NumSpeel --> %v \n%T \n", numSpeel, numSpeel)
	fmt.Println("numSpeel", len(numSpeel))
	fmt.Println("Heare you will see the magic of map")
	fmt.Println("numSpeel[2] --> ", numSpeel[2])

	var myNumber = 10
	var myPtr = &myNumber

	fmt.Println("My Pointer Address: ", myPtr)
	fmt.Println("My Pointer value: ", *myPtr)

	*myPtr = *myPtr + 2
	fmt.Println("addition Using the: ", myNumber)
	fmt.Println("Address of My Pointer Address: ", myPtr)

	//refrance means & sign
	//*ptr means the what inside in the ptr

	/*var ptr *int
	  fmt.Println("default value ", ptr)
	*/
}

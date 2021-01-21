package main //Default

import "fmt" //Deafult

//Main function
func main() {
	var number int = 10                            // Declaretion of an variable
	fmt.Println("Number: ", number)                // Printing variable
	fmt.Printf("Type of Number: %T", number)       //printing type of number
	fmt.Println()                                  //For new line
	fmt.Printf("Number -> Hexadecimal %X", number) //hexadecimal
	fmt.Println()
	fmt.Printf("Number -> Octal %o ", number) //Octal
	fmt.Println()
	fmt.Printf("Number -> Decimal %d ", number) //Decimal
	fmt.Println()
	fmt.Printf("Number -> Binary %b", number) //Binary
	fmt.Println()
}

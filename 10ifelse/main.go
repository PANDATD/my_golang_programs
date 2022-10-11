package main

import "fmt"

func main() {
	fmt.Println(" if else if ")
	FactCheck(18)

}

func FactCheck(age int) int {
	if age == 18 {
		fmt.Println("Your Adult ")
	} else if age < 18 {
		fmt.Println("Your Kid Now!")
	} else {
		fmt.Println("Your im looking for...")
	}
	return 0
}


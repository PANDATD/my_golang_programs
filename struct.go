package main

// imported required package

import "fmt"
// Heare i declared Struct of coddersclub

type coddersclub struct {
	name     string
	age      int
	position string
}

// This var syntax is user for multipal varibale Declaretion
var (
	member1 = coddersclub{
		"Mr. Tejas Dixit",
		19,
		"Founder",
	}

	member2 = coddersclub{
		"Mr. Datta Kale",
		21,
		"CEO",
	}

	member3 = coddersclub{
		"Mr. Kunal Jagate",
		20,
		"CSO",
	}

	member4 = coddersclub{
		"Mr. Vignesh Gawali",
		19,
		"CTO",
	}
)

func main() {
	// execution of funcation or code
	fmt.Println(" Stuct of CodderscluB ")

	fmt.Println(member1, member2, member3, member4)
}

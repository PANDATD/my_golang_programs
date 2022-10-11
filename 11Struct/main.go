package main

import "fmt"

func main() {

	fmt.Println("Hello there now we are looking Structs")

	tejas := User{"Tejas", 20, true}

	fmt.Println("Tejas: ", tejas)
	fmt.Printf("Tejas: %+v\n", tejas)
	fmt.Printf("Name:%v Age:%v\n", tejas.Name, tejas.Age)

	vbg := User{"Vignesh", 20, true}
	fmt.Print(vbg)

}

type User struct {
	Name   string
	Age    int
	Status bool
}

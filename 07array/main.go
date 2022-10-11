package main	

import (
	"fmt"
)

func main(){

	fmt.Println("welcome to array in golan")

	var fruitlist [4]string

	fruitlist[0]= "Apple"
	fruitlist[1]= "tomato"
	fruitlist[3]= "mango"

	fmt.Println("Fruit lis is: ", fruitlist)
	fmt.Println("fruit list is: ", len(fruitlist))

	
}

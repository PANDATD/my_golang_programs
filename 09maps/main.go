package main 

import "fmt"

func main (){
	
	fmt.Println("Maps in goloan")
	languges := make(map[string]string)

	languges["js"] = "JavaScripts"
	languges["py"] = "Python"
	languges["rb"] = "Ruby"


	fmt.Println("List of languges", languges)
	fmt.Println("js shorts for ", languges["js"])


	for key, value := range languges{
		fmt.Printf("KEY: %v VALUE: %v \n", key, value)
	}
}
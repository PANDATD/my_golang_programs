package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Slices --> ")

	var fruitlist = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mangos", "banana")

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 235
	highScore[2] = 236
	highScore[3] = 237
	//highScore[4]=238

	highScore = append(highScore, 555, 666, 321)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//How ro remove a value from Sliced based on index

	var coures = []string{"react", "js", "python", "swift"}
	fmt.Println(coures)

	var index int = 2
	coures = append(coures[:index], coures[index+1:]...)
	fmt.Print(coures)

}

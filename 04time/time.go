package main 

import (
	"fmt"
	"time"
)


func main(){

	fmt.Println("Time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.August, 10,23,23,0,0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}	

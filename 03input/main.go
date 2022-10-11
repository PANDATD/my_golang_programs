package main 

import "fmt"
import "bufio"
import "os"

func main(){


fmt.Println("Enter Your name: ")

reader:= bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
if err != nil{
	panic(err)
}
fmt.Println("Name is: ",input)

}